package files

import (
	"fmt"
	"os"
)

func CreateFileIfIsnotExist(filepath string) (*os.File, error) {

	if _, err := os.Stat(filepath); err == nil {
		return nil, fmt.Errorf("файл уже существует: %s", filepath)
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("ошибка при проверке существования файла: %v", err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return nil, fmt.Errorf("ошибка при создании файла: %v", err)
	}

	return file, nil
}
