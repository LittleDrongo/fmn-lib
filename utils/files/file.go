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
		return nil, fmt.Errorf("ошибка при проверке существования файла: %v", err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании файла: %v", err)
	}

	return file, nil
}
