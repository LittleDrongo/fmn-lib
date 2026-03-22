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
First create an instance of the struct that will be populated.

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
