package update

import (
	"bytes"
	"io"
	"net/http"
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
    {"name": "mallard_0.3.0_windows_amd64.zip",   "browser_download_url": "https://example.com/windows_amd64"}
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
	executablePath = func() (string, error) {
		return "/opt/homebrew/Cellar/mallard/0.2.0/bin/mallard", nil
	}
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
	// A path outside any package manager so the guard does not fire.
	executablePath = func() (string, error) { return t.TempDir() + "/mallard", nil }
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
