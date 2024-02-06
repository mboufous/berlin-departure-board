package bvg

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func loadMockedBVGResponse(t *testing.T, filename string) io.ReadCloser {
	t.Helper()
	path := filepath.Join("testdata", filename)
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", path, err)
	}
	return io.NopCloser(bytes.NewReader(content))
}
