package files

import (
	"os"
	"path/filepath"

	"github.com/LittleDrongo/fmn-lib/exception"
)

func MakeDirIfIsNotExist(path string) error {

	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		return exception.DropUp(err, "Ошибка при создании папки:")
	}

	return nil

}
