package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LittleDrongo/go_libs/utils/files"

	"github.com/LittleDrongo/go_libs/errors"
)

type Writer interface {
	Write(interface{}, string)
}

func Write(data interface{}, filepath string) {

	files.MakeDirIfIsNotExist(filepath)

	file, err := json.MarshalIndent(data, "", "	")
	errors.Println(err, "Ошибка при создании объекта данных JSON")

	err = os.WriteFile(filepath, file, 0644)
	errors.Println(err, "Ошибка сохранения файла JSON")
}

func Print(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	errors.Fatalln(err, "Ошибка при создании объекта данных JSON:")
	fmt.Println(string(jsonData))
}
