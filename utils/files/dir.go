package files

import (
	"os"
	"path/filepath"
)

func MakeDirIfIsNotExist(path string) error {

	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return nil

}
