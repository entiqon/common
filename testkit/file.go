package testkit

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var updateFiles = flag.Bool("update-files", false, "rewrite expected files with current output")

// WriteFiles creates a temp dir holding the given files (slash-separated
// relative path → content) and returns its path. It is removed when the test
// ends.
//
//	dir := testkit.WriteFiles(t, map[string]string{
//		"config.yaml":      "port: 8080",
//		"migrations/1.sql": "CREATE TABLE leads (...);",
//	})
func WriteFiles(
	t testing.TB,
	files map[string]string,
) string {
	t.Helper()

	dir := t.TempDir()
	for name, content := range files {
		path := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("WriteFiles: %v", err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("WriteFiles: %v", err)
		}
	}

	return dir
}

// ReadFile returns a file's contents, failing the test immediately on error.
func ReadFile(
	t testing.TB,
	path string,
) string {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	return string(b)
}

// EqualFile reports when got differs from testdata/<name> in the calling
// test's package. Run `go test -update-files` (or UPDATE_FILES=1) to create
// or refresh the file.
//
//	testkit.EqualFile(t, "export/leads.csv", out)
func EqualFile(
	t testing.TB,
	name string,
	got []byte,
) {
	t.Helper()

	path := filepath.Join("testdata", filepath.FromSlash(name))

	if *updateFiles || os.Getenv("UPDATE_FILES") == "1" {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("EqualFile: %v", err)
		}
		if err := os.WriteFile(path, got, 0o644); err != nil {
			t.Fatalf("EqualFile: %v", err)
		}
		return
	}

	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("EqualFile: %v (run with -update-files to create it)", err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("output differs from %s (-expected +got):\n%s\nrun with -update-files to accept",
			path, Diff(string(want), string(got)))
	}
}
