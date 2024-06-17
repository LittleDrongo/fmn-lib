package files

import (
	"os"
)

func MakeDirIfIsNotExist(path string) error {

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil

}
