// Package update implements mallard's binary self-update ("mallard upgrade").
//
// It queries the GitHub releases API for the latest tag, compares it against
// the running version, and — when newer — downloads the matching release asset,
// verifies its SHA-256 against checksums.txt, extracts the mallard binary, and
// atomically replaces the running executable.
//
// The package is dependency-free (stdlib only). The HTTP client and the
// "current executable path" resolver are package-level vars so tests can inject
// canned responses without touching the network or the filesystem.
package update

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// repoOwner/repoName identify the GitHub repository releases are fetched from.
const (
	repoOwner = "educlopez"
	repoName  = "mallard"
)

// releasesURL is the GitHub API endpoint for the latest mallard release.
const releasesURL = "https://api.github.com/repos/" + repoOwner + "/" + repoName + "/releases/latest"

// Body-read limits to cap memory usage and guard against malicious/corrupt responses.
const (
	maxAPIBodyBytes     = 10 * 1024 * 1024  // 10 MB — GitHub releases JSON
	maxArchiveBytes     = 200 * 1024 * 1024 // 200 MB — release archive download
	maxChecksumBytes    = 1 * 1024 * 1024   // 1 MB — checksums.txt
	maxBinaryEntryBytes = 200 * 1024 * 1024 // 200 MB — extracted binary entry
)

// Injection points for testing. Tests override these to avoid the network and
// to control which executable path the self-replace logic sees.
var (
	// httpClient is the HTTP client used for all outbound requests.
	// It carries a generous timeout so a stalled server does not hang the CLI
	// indefinitely; tests replace it with a mock transport.
	httpClient = &http.Client{Timeout: 120 * time.Second}
	// executablePath resolves the path of the currently running binary.
	executablePath = os.Executable
)

// release is the subset of the GitHub releases API payload we need.
type release struct {
	TagName string  `json:"tag_name"`
	Assets  []asset `json:"assets"`
}

type asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

// Options configures an upgrade run.
type Options struct {
	// CurrentVersion is the running binary's version (main.version). "dev" is
	// treated as always-older.
	CurrentVersion string
	// Force replaces even when not strictly newer (and lets "dev" self-replace).
	Force bool
	// Check only reports the available version; nothing is downloaded.
	Check bool
	// DryRun resolves and reports the action but writes nothing to disk.
	DryRun bool
}

// Run performs the upgrade (or check) and writes a human report to w.
func Run(w io.Writer, opts Options) error {
	// Env guard.
	if os.Getenv("MALLARD_NO_SELF_UPDATE") == "1" {
		fmt.Fprintln(w, "  Self-update disabled via MALLARD_NO_SELF_UPDATE=1.")
		return nil
	}

	rel, err := fetchLatest()
	if err != nil {
		return fmt.Errorf("fetch latest release: %w", err)
	}
	latest := rel.TagName

	cmp := compareVersions(opts.CurrentVersion, latest)
	fmt.Fprintf(w, "  Current: %s\n  Latest:  %s\n", opts.CurrentVersion, latest)

	switch {
	case cmp == 0 && !opts.Force:
		fmt.Fprintln(w, "  Already up to date.")
		return nil
	case cmp > 0 && !opts.Force:
		fmt.Fprintln(w, "  Current version is newer than the latest release; nothing to do.")
		return nil
	}

	// At this point we are behind (cmp < 0), or --force was given. For "dev"
	// without --force we print a notice and skip the actual replacement.
	if isDev(opts.CurrentVersion) && !opts.Force {
		fmt.Fprintln(w, "  Running a dev build — skipping self-replace. Use --force to overwrite with the latest release.")
		return nil
	}

	if opts.Check {
		fmt.Fprintf(w, "  Update available: %s -> %s (run `mallard upgrade` to apply).\n", opts.CurrentVersion, latest)
		return nil
	}

	exe, err := executablePath()
	if err != nil {
		return fmt.Errorf("resolve current executable: %w", err)
	}
	resolved, err := filepath.EvalSymlinks(exe)
	if err != nil {
		return fmt.Errorf("resolve symlinks for %s: %w", exe, err)
	}
	exe = resolved

	// Package-manager guard: never self-replace a brew/scoop-managed binary.
	if mgr := detectPackageManager(exe); mgr != "" {
		printPackageManagerHint(w, mgr)
		return nil
	}

	a, err := selectAsset(rel.Assets, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	if opts.DryRun {
		fmt.Fprintf(w, "  (dry run) would download %s\n  (dry run) would replace %s\n", a.Name, exe)
		return nil
	}

	if err := downloadAndReplace(w, rel, a, exe); err != nil {
		return err
	}
	fmt.Fprintf(w, "  Upgraded to %s.\n", latest)
	return nil
}

// fetchLatest queries the GitHub releases API via the injectable client.
func fetchLatest() (release, error) {
	req, err := http.NewRequest(http.MethodGet, releasesURL, nil)
	if err != nil {
		return release{}, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "mallard-upgrade")

	resp, err := httpClient.Do(req)
	if err != nil {
		return release{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return release{}, fmt.Errorf("github API returned status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxAPIBodyBytes))
	if err != nil {
		return release{}, err
	}
	var rel release
	if err := json.Unmarshal(body, &rel); err != nil {
		return release{}, err
	}
	if rel.TagName == "" {
		return release{}, fmt.Errorf("github API returned a release with no tag_name")
	}
	return rel, nil
}

// isDev reports whether v is a development build (never a real release).
func isDev(v string) bool {
	v = strings.TrimSpace(v)
	return v == "" || v == "dev"
}

// compareVersions compares two semantic-ish versions, tolerating a leading "v".
// It returns -1 if a<b, 0 if equal, +1 if a>b. A "dev" (or empty) a is always
// older than any real release (returns -1).
func compareVersions(a, b string) int {
	if isDev(a) {
		return -1
	}
	if isDev(b) {
		return 1
	}
	pa := parseVersion(a)
	pb := parseVersion(b)
	for i := 0; i < len(pa) || i < len(pb); i++ {
		var x, y int
		if i < len(pa) {
			x = pa[i]
		}
		if i < len(pb) {
			y = pb[i]
		}
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}
	return 0
}

// parseVersion turns "v1.2.3" / "1.2.3" into [1 2 3]. Any pre-release suffix
// (e.g. "-rc1") on a segment is dropped; non-numeric segments become 0.
func parseVersion(v string) []int {
	v = strings.TrimSpace(v)
	v = strings.TrimPrefix(v, "v")
	// Drop a build/prerelease suffix on the whole string.
	if i := strings.IndexAny(v, "-+"); i >= 0 {
		v = v[:i]
	}
	parts := strings.Split(v, ".")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			n = 0
		}
		out = append(out, n)
	}
	return out
}

// detectPackageManager returns "brew", "scoop", or "" based on the executable
// path. Homebrew installs land under .../Cellar/mallard/... (or a linked opt
// path); Scoop installs under .../scoop/apps/mallard/....
func detectPackageManager(exe string) string {
	// Normalize backslashes to forward slashes regardless of host OS so the
	// substring checks work for Windows paths even when tests run on unix.
	p := strings.ReplaceAll(exe, "\\", "/")
	lower := strings.ToLower(p)
	if strings.Contains(p, "/Cellar/mallard/") ||
		strings.Contains(p, "/opt/mallard/") ||
		strings.Contains(p, "/Homebrew/") ||
		strings.Contains(lower, "/homebrew/") {
		return "brew"
	}
	if strings.Contains(lower, "/scoop/apps/mallard/") {
		return "scoop"
	}
	return ""
}

func printPackageManagerHint(w io.Writer, mgr string) {
	switch mgr {
	case "brew":
		fmt.Fprintln(w, "  Installed via Homebrew — run `brew upgrade mallard` instead.")
	case "scoop":
		fmt.Fprintln(w, "  Installed via Scoop — run `scoop update mallard` instead.")
	}
}

// selectAsset picks the release asset matching the given OS/arch. The names
// follow the .goreleaser.yaml template:
//
//	mallard_<version>_<os>_<arch>.tar.gz   (.zip on windows)
//
// where <version> has no leading "v". We match on the os/arch infix plus the
// expected extension rather than reconstructing the version, so it is robust to
// version-string formatting.
func selectAsset(assets []asset, goos, goarch string) (asset, error) {
	ext := ".tar.gz"
	if goos == "windows" {
		ext = ".zip"
	}
	infix := "_" + goos + "_" + goarch + ext
	for _, a := range assets {
		if strings.HasSuffix(a.Name, infix) {
			return a, nil
		}
	}
	return asset{}, fmt.Errorf("no release asset found for %s/%s (looked for *%s)", goos, goarch, infix)
}

// checksumAssetURL returns the download URL for checksums.txt by looking for
// it in the release asset list, or falls back to a predictable URL pattern.
func checksumAssetURL(assets []asset, tagName string) string {
	for _, a := range assets {
		if a.Name == "checksums.txt" {
			return a.BrowserDownloadURL
		}
	}
	// Fallback: construct the URL from the known pattern.
	return fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/checksums.txt", repoOwner, repoName, tagName)
}

// verifyChecksum downloads checksums.txt from the release, finds the entry for
// archiveName, and verifies that the SHA-256 of archiveData matches. It returns
// an error if checksums.txt cannot be fetched, the entry is missing, or the
// digest does not match.
func verifyChecksum(assets []asset, tagName, archiveName string, archiveData []byte) error {
	url := checksumAssetURL(assets, tagName)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build checksums.txt request: %w", err)
	}
	req.Header.Set("User-Agent", "mallard-upgrade")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("download checksums.txt: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("checksums.txt download returned status %d", resp.StatusCode)
	}

	body, err := readCapped(resp.Body, maxChecksumBytes)
	if err != nil {
		return fmt.Errorf("read checksums.txt: %w", err)
	}

	// Parse the goreleaser checksum format: "<hex>  <filename>" (two spaces).
	var expected string
	for _, line := range strings.Split(string(body), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Accept both single-space and double-space separators.
		parts := strings.Fields(line)
		if len(parts) == 2 && parts[1] == archiveName {
			expected = parts[0]
			break
		}
	}
	if expected == "" {
		return fmt.Errorf("archive %q not found in checksums.txt — refusing to install unverified binary", archiveName)
	}

	sum := sha256.Sum256(archiveData)
	actual := hex.EncodeToString(sum[:])
	if actual != expected {
		return fmt.Errorf("checksum mismatch for %s:\n  expected: %s\n  got:      %s", archiveName, expected, actual)
	}
	return nil
}

// downloadAndReplace downloads the asset, verifies its checksum, extracts the
// mallard binary, writes it to a temp file beside exe, chmod +x, and atomically
// renames it over exe.
func downloadAndReplace(w io.Writer, rel release, a asset, exe string) error {
	fmt.Fprintf(w, "  Downloading %s ...\n", a.Name)
	data, err := download(a.BrowserDownloadURL, maxArchiveBytes)
	if err != nil {
		return fmt.Errorf("download asset: %w", err)
	}

	fmt.Fprintf(w, "  Verifying checksum ...\n")
	if err := verifyChecksum(rel.Assets, rel.TagName, a.Name, data); err != nil {
		return fmt.Errorf("checksum verification: %w", err)
	}

	bin, err := extractBinary(a.Name, data)
	if err != nil {
		return fmt.Errorf("extract binary: %w", err)
	}

	dir := filepath.Dir(exe)
	tmp, err := os.CreateTemp(dir, ".mallard-upgrade-*")
	if err != nil {
		return fmt.Errorf("create temp file next to executable (check permissions on %s): %w", dir, err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName) // no-op if the rename succeeded

	if _, err := tmp.Write(bin); err != nil {
		tmp.Close()
		return fmt.Errorf("write new binary: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close new binary: %w", err)
	}
	if err := os.Chmod(tmpName, 0o755); err != nil {
		return fmt.Errorf("chmod new binary: %w", err)
	}
	if err := os.Rename(tmpName, exe); err != nil {
		return fmt.Errorf("replace %s (insufficient permissions? try with elevated privileges or your package manager): %w", exe, err)
	}
	return nil
}

// readCapped reads at most limit bytes from r and returns an explicit error if
// the stream holds more, instead of silently truncating the data.
func readCapped(r io.Reader, limit int64) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(r, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(data)) > limit {
		return nil, fmt.Errorf("response exceeds %d byte limit", limit)
	}
	return data, nil
}

// download fetches the URL bytes via the injectable client, capped at limit bytes.
func download(url string, limit int64) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "mallard-upgrade")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download returned status %d", resp.StatusCode)
	}
	return readCapped(resp.Body, limit)
}

// extractBinary pulls the mallard executable out of a .tar.gz or .zip archive.
func extractBinary(assetName string, data []byte) ([]byte, error) {
	if strings.HasSuffix(assetName, ".zip") {
		return extractFromZip(data)
	}
	return extractFromTarGz(data)
}

// binaryName is the name of the binary inside the archive (mallard.exe on
// windows, mallard elsewhere).
func binaryName() string {
	if runtime.GOOS == "windows" {
		return "mallard.exe"
	}
	return "mallard"
}

func extractFromTarGz(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	want := binaryName()
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if filepath.Base(hdr.Name) == want {
			return readCapped(tr, maxBinaryEntryBytes)
		}
	}
	return nil, fmt.Errorf("binary %q not found in archive", want)
}

func extractFromZip(data []byte) ([]byte, error) {
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	want := binaryName()
	for _, f := range zr.File {
		if filepath.Base(f.Name) == want {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return readCapped(rc, maxBinaryEntryBytes)
		}
	}
	return nil, fmt.Errorf("binary %q not found in archive", want)
}
