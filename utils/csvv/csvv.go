package csvv

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

func sample() {
	type MyStruct struct {
		Id           string    `csv:"key"`
		FieldString  string    `csv:"FieldString"`
		FieldFloat64 float64   `csv:"FieldFloat64"`
		FieldInt     int       `csv:"FieldInt"`
		FieldTime    time.Time `csv:"FieldTime"`
	}

	type Repository struct {
		MyStructs map[string]MyStruct
	}

	// repo := Repository{
	// 	MyStructs: make(map[string]MyStruct),
	// }

	repo := make(map[string]MyStruct)

	// Импорт данных из CSV
	err := ImportMapFromCSV("data.csv", &repo)
	if err != nil {
		fmt.Println("Error importing CSV:", err)
		return
	}

	// Дополняем данные
	repo["800"] = MyStruct{
		Id:           "800",
		FieldString:  "str3",
		FieldFloat64: 64.4,
		FieldInt:     30,
		FieldTime:    time.Now(),
	}

	for key, value := range repo {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}

	// Экспорт данных в CSV
	err = ExportMapToCSV("data.csv", repo)
	if err != nil {
		fmt.Println("Error exporting CSV:", err)
		return
	}

	fmt.Println("CSV import/export completed successfully")
}

/*
Сначала создаётся map который будет заполняться

	repo := make(map[string]MyStruct)
	err := ImportMapFromCSV("data.csv", &repo)
*/
func ImportMapFromCSV(filePath string, data interface{}) error {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		dataValue := reflect.ValueOf(data)
		if dataValue.Kind() != reflect.Ptr || dataValue.Elem().Kind() != reflect.Map {
			return fmt.Errorf("data should be a pointer to a map")
		}
		newMap := reflect.MakeMap(dataValue.Elem().Type())
		dataValue.Elem().Set(newMap)
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Ptr || dataValue.Elem().Kind() != reflect.Map {
		return fmt.Errorf("data should be a pointer to a map")
	}

	mapType := dataValue.Elem().Type()
	keyType := mapType.Key()
	elemType := mapType.Elem()

	var keyFieldName string
	for i := 0; i < elemType.NumField(); i++ {
		if tag := elemType.Field(i).Tag.Get("csv"); tag == "key" {
			keyFieldName = elemType.Field(i).Name
			break
		}
	}

	if keyFieldName == "" {
		return fmt.Errorf("no field with csv:\"key\" tag found")
	}

	newMap := reflect.MakeMap(mapType)

	for _, record := range records {
		elem := reflect.New(elemType).Elem()
		var key reflect.Value

		for i, header := range headers {
			field := elem.FieldByName(header)
			if !field.IsValid() {
				continue
			}
			fieldType := field.Type()
			fieldValue := record[i]

			switch fieldType.Kind() {
			case reflect.String:
				field.SetString(fieldValue)
			case reflect.Float64:
				floatVal, err := strconv.ParseFloat(fieldValue, 64)
				if err != nil {
					return err
				}
				field.SetFloat(floatVal)
			case reflect.Int:
				intVal, err := strconv.Atoi(fieldValue)
				if err != nil {
					return err
				}
				field.SetInt(int64(intVal))
			case reflect.Struct:
				if fieldType == reflect.TypeOf(time.Time{}) {
					timeVal, err := time.Parse(time.RFC3339, fieldValue)
					if err != nil {
						return err
					}
					field.Set(reflect.ValueOf(timeVal))
				}
			default:
				return fmt.Errorf("unsupported field type: %s", fieldType.Kind())
			}

			if header == keyFieldName {
				switch keyType.Kind() {
				case reflect.String:
					key = reflect.ValueOf(fieldValue)
				case reflect.Int:
					intVal, err := strconv.Atoi(fieldValue)
					if err != nil {
						return err
					}
					key = reflect.ValueOf(intVal)
				default:
					return fmt.Errorf("unsupported key type: %s", keyType.Kind())
				}
			}
		}

		if !key.IsValid() {
			return fmt.Errorf("key field %s is not found in record", keyFieldName)
		}

		newMap.SetMapIndex(key, elem)
	}

	dataValue.Elem().Set(newMap)

	return nil
}
func ExportMapToCSV(filePath string, data interface{}) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Map {
		return fmt.Errorf("data should be a map")
	}

	elemType := dataValue.Type().Elem()

	numFields := elemType.NumField()
	headers := make([]string, numFields)
	var keyFieldName string
	for i := 0; i < numFields; i++ {
		headers[i] = elemType.Field(i).Name
		if tag := elemType.Field(i).Tag.Get("csv"); tag == "key" {
			keyFieldName = elemType.Field(i).Name
		}
	}

	if keyFieldName == "" {
		return fmt.Errorf("no field with csv:\"key\" tag found")
	}

	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, key := range dataValue.MapKeys() {
		elem := dataValue.MapIndex(key)

		record := make([]string, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			fieldType := field.Type()

			switch fieldType.Kind() {
			case reflect.String:
				record[i] = field.String()
			case reflect.Float64:
				record[i] = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Int:
				record[i] = strconv.Itoa(int(field.Int()))
			case reflect.Struct:
				if fieldType == reflect.TypeOf(time.Time{}) {
					record[i] = field.Interface().(time.Time).Format(time.RFC3339)
				}
			default:
				return fmt.Errorf("unsupported field type: %s", fieldType.Kind())
			}

			if elemType.Field(i).Name == keyFieldName {
				switch key.Kind() {
				case reflect.String:
					record[i] = key.String()
				case reflect.Int:
					record[i] = strconv.Itoa(int(key.Int()))
				default:
					return fmt.Errorf("unsupported key type: %s", key.Kind())
				}
			}
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
