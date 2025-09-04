package gobs

import (
	"encoding/gob"
	"os"
)

func Export(data any, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

/*
Сначала создаётся экземпляр класса который будет заполняться

	var myStrc myStruct
	gobs.Import("data/file.job", &myStrc)
*/
func Import(filepath string, data any) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(data)
}
