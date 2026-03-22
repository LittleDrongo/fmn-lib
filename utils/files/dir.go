package files

import (
	"os"
	"path/filepath"
)

// EnsureDir creates a directory and all missing parents if needed.
func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// EnsureDirForFile creates the parent directory for a file path.
func EnsureDirForFile(path string) error {
	return EnsureDir(filepath.Dir(path))
}
