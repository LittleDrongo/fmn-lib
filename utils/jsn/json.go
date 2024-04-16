package jsn

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LittleDrongo/fmn-lib/exception"
	"github.com/LittleDrongo/fmn-lib/utils/files"
)

func Export(data interface{}, filepath string) error {

	files.MakeDirIfIsNotExist(filepath)

	file, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return exception.DropUp(err, "Ошибка при создании объекта данных JSON:")
	}

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		return exception.DropUp(err, "Ошибка сохранения файла JSON:")
	}

	return nil
}

func Print(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return exception.DropUp(err, "Ошибка при создании объекта данных JSON:")
	}
	fmt.Println(string(jsonData))
	return nil
}

// 1. Сначала создаётся экземпляр класса который будет заполняться: var myStrc MySturct
//
// 2. В аргументах Import("filepath", &myStrc) передаётся указатель переменной для заполнения данными из файла
func Import(filepath string, anyTypePointer interface{}) error {

	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &anyTypePointer)
	if err != nil {
		return err
	}

	return nil
}
