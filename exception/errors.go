package exception

import (
	"errors"
	"fmt"
	"log"
)

func Print(err error, message ...interface{}) {
	if err != nil {
		log.Print(fmt.Sprintf("%v", message...), err)
	}

}

func Println(err error, message ...interface{}) {
	if err != nil {
		log.Println(fmt.Sprintf("%v", message...), err)
	}
}

func Fatal(err error, message ...interface{}) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%v", message...), err)
	}

}

func Fatalln(err error, message ...interface{}) {
	if err != nil {
		log.Fatalln(fmt.Sprintf("%v", message...), err)
	}
}

func DropUp(err error, message ...interface{}) error {
	str := fmt.Sprint(fmt.Sprintf("%v", message...), " ", err)
	if err != nil {
		return errors.New(str)
	}
	return nil
}
