package gobs

import (
	"encoding/gob"
	"os"
)

func Export(data interface{}, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

/*
Пример импорта файла
var myobj someStruct{}

gobs.Import("file.gob", &data)
*/
func Import(filepath string, data interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(data)
}
