package files

import (
	"os"
	"path/filepath"

	"github.com/LittleDrongo/go_libs/errors"
)

func MakeDirIfIsNotExist(path string) {

	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		errors.Println(err, "Ошибка при создании папки:")
	}

}
