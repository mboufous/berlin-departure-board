package bvg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func loadMockedBVGResponse(filename string) (io.ReadCloser, error) {
	path := filepath.Join("testdata", filename)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", path, err)
	}
	return io.NopCloser(bytes.NewReader(content)), nil
}
