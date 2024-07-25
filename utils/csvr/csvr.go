package csvr

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func sampleCsv() {

	type MyStruct struct {
		Id       int `csv:"key"`
		FieldOne string
		FieldTwo int
	}
	filename := "data.csv"

	var err error

	// // Импортируем данные из CSV файла
	importedData := make(map[int]MyStruct)
	err = ImportMapFromCSV(filename, &importedData)
	if err != nil {
		fmt.Println("Error importing from CSV:", err)
		return
	}

	// Добавляем новые данные
	newRecord := MyStruct{Id: 101, FieldOne: "NewFieldOne", FieldTwo: 30}
	importedData[newRecord.Id] = newRecord

	// Экспортируем данные обратно в CSV файл
	err = ExportMapToCSV(importedData, filename)
	if err != nil {
		fmt.Println("Error exporting to CSV:", err)
		return
	}

	fmt.Println("Data successfully saved to", filename)
}

func ExportMapToCSV(data interface{}, filepath string) error {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Map {
		return errors.New("data must be a map")
	}

	if val.Len() == 0 {
		return errors.New("data map is empty")
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	var headers []string
	for _, key := range val.MapKeys() {
		elem := val.MapIndex(key)
		elemType := elem.Type()
		for i := 0; i < elemType.NumField(); i++ {
			headers = append(headers, elemType.Field(i).Name)
		}
		break
	}
	writer.Write(headers)

	// Write data
	for _, key := range val.MapKeys() {
		elem := val.MapIndex(key)
		var record []string
		for i := 0; i < elem.NumField(); i++ {
			record = append(record, fmt.Sprintf("%v", elem.Field(i).Interface()))
		}
		writer.Write(record)
	}

	return nil
}

func ImportMapFromCSV(filepath string, data interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) < 2 {
		return errors.New("CSV file must contain at least one record")
	}

	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Map {
		return errors.New("data must be a pointer to a map")
	}

	valElem := val.Elem()
	keyType := valElem.Type().Key()
	elemType := valElem.Type().Elem()

	headers := records[0]
	for _, record := range records[1:] {
		newElem := reflect.New(elemType).Elem()
		var key reflect.Value
		for i, header := range headers {
			field := newElem.FieldByName(header)
			if !field.IsValid() {
				return fmt.Errorf("no such field: %s in obj", header)
			}
			value := record[i]
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				field.SetInt(intValue)
			case reflect.Float32, reflect.Float64:
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				field.SetFloat(floatValue)
			default:
				return fmt.Errorf("unsupported kind: %s", field.Kind())
			}
			// Set key if field has the tag `csv:"key"`
			if elemType.Field(i).Tag.Get("csv") == "key" {
				key = reflect.ValueOf(field.Interface()).Convert(keyType)
			}
		}
		if !key.IsValid() {
			return errors.New("no key field found in struct")
		}
		valElem.SetMapIndex(key, newElem)
	}

	return nil
}

func CreateCSVWithHeaders(structType interface{}, filepath string) error {
	if _, err := os.Stat(filepath); err == nil {
		// Файл существует
		return nil
	} else if !os.IsNotExist(err) {
		// Ошибка при попытке проверить наличие файла
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	val := reflect.ValueOf(structType)
	if val.Kind() != reflect.Struct {
		return errors.New("structType must be a struct")
	}

	var headers []string
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Name)
	}

	writer.Write(headers)

	return nil
}
