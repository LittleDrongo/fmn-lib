package yamlt

import (
	"fmt"
	"os"

	"github.com/LittleDrongo/fmn-lib/exception"
	"github.com/LittleDrongo/fmn-lib/utils/files"
	"gopkg.in/yaml.v2"
)

// Метод экспортирует любую структуру в формате YAML файла.
func Export(data interface{}, filepath string) error {
	files.MakeDirIfIsNotExist(filepath)

	file, err := yaml.Marshal(data)
	if err != nil {
		return exception.DropUp(err, "Ошибка при создании объекта данных YAML:")
	}

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		return exception.DropUp(err, "Ошибка сохранения файла YAML:")
	}

	return nil
}

// Метод печатать любую структуру в формате YAML файла.
func Print(data interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return exception.DropUp(err, "ошибка при создании объекта данных YAML:")
	}
	fmt.Println(string(yamlData))
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

	err = yaml.Unmarshal(file, anyTypePointer)
	if err != nil {
		return err
	}

	return nil
}
