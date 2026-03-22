package files

import (
	"fmt"
	"os"
)

func CreateFileIfIsnotExist(filepath string) (*os.File, error) {

	if info, err := os.Stat(filepath); err == nil {
		if !info.IsDir() {
			file, err := os.Open(filepath)
			if err != nil {
				return nil, err
			}
			return file, nil
		}

	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check file existence: %v", err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}

	return file, nil
}
