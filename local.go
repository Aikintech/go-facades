package gofacades

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Local struct{}

func (l *Local) WriteFile(path string, data []byte, permissions uint32) error {
	if permissions == 0 {
		permissions = 0644
	}

	// Extract directory path
	dir := filepath.Dir(path)

	// Create directories if they don't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return errors.Wrap(err, "while creating directories")
	}

	// Write the file
	if err := os.WriteFile(path, data, os.FileMode(permissions)); err != nil {
		return errors.Wrap(err, "while writing file")
	}
	return nil
}
