package update

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

// roundTripFunc lets a func stand in for an http.RoundTripper.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// withMockClient installs a client that returns body for the releases URL and
// restores the original client when the test ends.
func withMockClient(t *testing.T, status int, body string) {
	t.Helper()
	orig := httpClient
	t.Cleanup(func() { httpClient = orig })
	httpClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: status,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
			}, nil
		}),
	}
}


const cannedRelease = `{
  "tag_name": "v0.3.0",
  "assets": [
    {"name": "mallard_0.3.0_darwin_arm64.tar.gz", "browser_download_url": "https://example.com/darwin_arm64"},
    {"name": "mallard_0.3.0_darwin_amd64.tar.gz", "browser_download_url": "https://example.com/darwin_amd64"},
    {"name": "mallard_0.3.0_linux_amd64.tar.gz",  "browser_download_url": "https://example.com/linux_amd64"},
    {"name": "mallard_0.3.0_windows_amd64.zip",   "browser_download_url": "https://example.com/windows_amd64"},
    {"name": "checksums.txt",                      "browser_download_url": "https://example.com/checksums.txt"}
  ]
}`

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"v0.2.0", "v0.3.0", -1},
		{"0.3.0", "0.3.0", 0},
		{"v0.3.0", "v0.2.0", 1},
		{"v1.0.0", "v0.9.9", 1},
		{"dev", "v0.3.0", -1},
		{"", "v0.3.0", -1},
		{"v0.3.0", "dev", 1},
		{"v0.3.0-rc1", "v0.3.0", 0}, // prerelease suffix dropped
		{"v0.3", "v0.3.0", 0},       // missing segment == 0
	}
	for _, tt := range tests {
		if got := compareVersions(tt.a, tt.b); got != tt.want {
			t.Errorf("compareVersions(%q,%q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestDetectPackageManager(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/opt/homebrew/Cellar/mallard/0.3.0/bin/mallard", "brew"},
		{"/usr/local/Cellar/mallard/0.3.0/bin/mallard", "brew"},
		{"/opt/homebrew/opt/mallard/bin/mallard", "brew"},
		{"C:\\Users\\me\\scoop\\apps\\mallard\\current\\mallard.exe", "scoop"},
		{"/Users/me/.local/bin/mallard", ""},
		{"/usr/local/bin/mallard", ""},
	}
	for _, tt := range tests {
		if got := detectPackageManager(tt.path); got != tt.want {
			t.Errorf("detectPackageManager(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestSelectAsset(t *testing.T) {
	rel, err := fetchTestRelease(t)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		goos, goarch string
		wantURL      string
		wantErr      bool
	}{
		{"darwin", "arm64", "https://example.com/darwin_arm64", false},
		{"darwin", "amd64", "https://example.com/darwin_amd64", false},
		{"linux", "amd64", "https://example.com/linux_amd64", false},
		{"windows", "amd64", "https://example.com/windows_amd64", false},
		{"linux", "arm64", "", true}, // not present in canned release
	}
	for _, tt := range tests {
		got, err := selectAsset(rel.Assets, tt.goos, tt.goarch)
		if tt.wantErr {
			if err == nil {
				t.Errorf("selectAsset(%s/%s) expected error", tt.goos, tt.goarch)
			}
			continue
		}
		if err != nil {
			t.Errorf("selectAsset(%s/%s) error: %v", tt.goos, tt.goarch, err)
			continue
		}
		if got.BrowserDownloadURL != tt.wantURL {
			t.Errorf("selectAsset(%s/%s) url = %q, want %q", tt.goos, tt.goarch, got.BrowserDownloadURL, tt.wantURL)
		}
	}
}

// fetchTestRelease parses the canned release JSON via fetchLatest with a mock.
func fetchTestRelease(t *testing.T) (release, error) {
	t.Helper()
	withMockClient(t, http.StatusOK, cannedRelease)
	return fetchLatest()
}

func TestRunCheckReportsAvailable(t *testing.T) {
	withMockClient(t, http.StatusOK, cannedRelease)
	var buf bytes.Buffer
	if err := Run(&buf, Options{CurrentVersion: "v0.2.0", Check: true}); err != nil {
		t.Fatalf("Run check: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "Update available") {
		t.Fatalf("expected 'Update available', got: %s", out)
	}
	if !strings.Contains(out, "v0.3.0") {
		t.Fatalf("expected latest version in output, got: %s", out)
	}
}

func TestRunAlreadyUpToDate(t *testing.T) {
	withMockClient(t, http.StatusOK, cannedRelease)
	var buf bytes.Buffer
	if err := Run(&buf, Options{CurrentVersion: "v0.3.0"}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !strings.Contains(buf.String(), "Already up to date") {
		t.Fatalf("expected up-to-date message, got: %s", buf.String())
	}
}

func TestRunDevSkipsWithoutForce(t *testing.T) {
	withMockClient(t, http.StatusOK, cannedRelease)
	var buf bytes.Buffer
	if err := Run(&buf, Options{CurrentVersion: "dev"}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !strings.Contains(buf.String(), "dev build") {
		t.Fatalf("expected dev-build notice, got: %s", buf.String())
	}
}

func TestRunNewerThanLatest(t *testing.T) {
	withMockClient(t, http.StatusOK, cannedRelease)
	var buf bytes.Buffer
	if err := Run(&buf, Options{CurrentVersion: "v9.9.9"}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !strings.Contains(buf.String(), "newer than the latest") {
		t.Fatalf("expected newer-than-latest message, got: %s", buf.String())
	}
}

func TestRunEnvGuard(t *testing.T) {
	t.Setenv("MALLARD_NO_SELF_UPDATE", "1")
	var buf bytes.Buffer
	if err := Run(&buf, Options{CurrentVersion: "v0.2.0"}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !strings.Contains(buf.String(), "disabled") {
		t.Fatalf("expected disabled message, got: %s", buf.String())
	}
}

func TestRunBrewGuard(t *testing.T) {
	withMockClient(t, http.StatusOK, cannedRelease)
	origExe := executablePath
	t.Cleanup(func() { executablePath = origExe })
	// Create a real temp directory whose path contains the brew Cellar pattern
	// so filepath.EvalSymlinks succeeds and detectPackageManager fires.
	brewDir := t.TempDir()
	// Rename temp dir to embed the cellar pattern in the path is not portable,
	// so instead we create a real file inside the temp dir and fake the path
	// by overriding executablePath to return a path we've arranged to exist.
	// We create a real file and craft the path reported by executablePath to
	// contain the brew marker: we put a "Cellar/mallard" subdirectory with the
	// binary inside, then return that path.
	binDir := brewDir + "/Cellar/mallard/0.2.0/bin"
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatalf("setup brew dir: %v", err)
	}
	binPath := binDir + "/mallard"
	if err := os.WriteFile(binPath, []byte("fake"), 0o755); err != nil {
		t.Fatalf("setup brew bin: %v", err)
	}
	executablePath = func() (string, error) { return binPath, nil }
	var buf bytes.Buffer
	// Behind latest and not a dev build, so it would normally try to replace —
	// the brew guard must intercept before any download.
	if err := Run(&buf, Options{CurrentVersion: "v0.2.0"}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if !strings.Contains(buf.String(), "brew upgrade mallard") {
		t.Fatalf("expected brew hint, got: %s", buf.String())
	}
}

func TestRunDryRunResolvesAsset(t *testing.T) {
	withMockClient(t, http.StatusOK, cannedRelease)
	origExe := executablePath
	t.Cleanup(func() { executablePath = origExe })
	// Create a real file so filepath.EvalSymlinks succeeds.
	// The path must be outside any package manager pattern so the guard does not fire.
	binPath := t.TempDir() + "/mallard"
	if err := os.WriteFile(binPath, []byte("fake"), 0o755); err != nil {
		t.Fatalf("setup bin: %v", err)
	}
	executablePath = func() (string, error) { return binPath, nil }
	var buf bytes.Buffer
	if err := Run(&buf, Options{CurrentVersion: "v0.2.0", DryRun: true}); err != nil {
		t.Fatalf("Run dry-run: %v", err)
	}
	if !strings.Contains(buf.String(), "would download") {
		t.Fatalf("expected dry-run download notice, got: %s", buf.String())
	}
}

func TestFetchLatestBadStatus(t *testing.T) {
	withMockClient(t, http.StatusNotFound, `{}`)
	if _, err := fetchLatest(); err == nil {
		t.Fatalf("expected error on 404")
	}
}

// ---------------------------------------------------------------------------
// Checksum verification tests
// ---------------------------------------------------------------------------

// sha256Hex returns the hex-encoded SHA-256 digest of data.
func sha256Hex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// cannedAssets returns a minimal asset list that includes checksums.txt pointing
// to the given URL.
func cannedAssetsWithChecksum(checksumURL string) []asset {
	return []asset{
		{Name: "mallard_0.3.0_linux_amd64.tar.gz", BrowserDownloadURL: "https://example.com/archive"},
		{Name: "checksums.txt", BrowserDownloadURL: checksumURL},
	}
}

func TestVerifyChecksumMatch(t *testing.T) {
	archiveData := []byte("fake-archive-content")
	archiveName := "mallard_0.3.0_linux_amd64.tar.gz"
	correctHex := sha256Hex(archiveData)

	checksumBody := fmt.Sprintf("%s  %s\n", correctHex, archiveName)
	assets := cannedAssetsWithChecksum("https://example.com/checksums.txt")

	orig := httpClient
	t.Cleanup(func() { httpClient = orig })
	httpClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(checksumBody)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	if err := verifyChecksum(assets, "v0.3.0", archiveName, archiveData); err != nil {
		t.Fatalf("expected no error on checksum match, got: %v", err)
	}
}

func TestVerifyChecksumMismatch(t *testing.T) {
	archiveData := []byte("fake-archive-content")
	archiveName := "mallard_0.3.0_linux_amd64.tar.gz"
	wrongHex := sha256Hex([]byte("different-content"))

	checksumBody := fmt.Sprintf("%s  %s\n", wrongHex, archiveName)
	assets := cannedAssetsWithChecksum("https://example.com/checksums.txt")

	orig := httpClient
	t.Cleanup(func() { httpClient = orig })
	httpClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(checksumBody)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	err := verifyChecksum(assets, "v0.3.0", archiveName, archiveData)
	if err == nil {
		t.Fatal("expected error on checksum mismatch, got nil")
	}
	if !strings.Contains(err.Error(), "checksum mismatch") {
		t.Fatalf("expected 'checksum mismatch' in error, got: %v", err)
	}
}

func TestVerifyChecksumMissingEntry(t *testing.T) {
	archiveData := []byte("fake-archive-content")
	archiveName := "mallard_0.3.0_linux_amd64.tar.gz"

	// checksums.txt lists a different archive — our entry is absent.
	checksumBody := "aabbcc  mallard_0.3.0_darwin_arm64.tar.gz\n"
	assets := cannedAssetsWithChecksum("https://example.com/checksums.txt")

	orig := httpClient
	t.Cleanup(func() { httpClient = orig })
	httpClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(checksumBody)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	err := verifyChecksum(assets, "v0.3.0", archiveName, archiveData)
	if err == nil {
		t.Fatal("expected error when archive entry is missing from checksums.txt, got nil")
	}
	if !strings.Contains(err.Error(), "not found in checksums.txt") {
		t.Fatalf("expected 'not found in checksums.txt' in error, got: %v", err)
	}
}
