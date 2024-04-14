package exception

import (
	"log"
)

func Print(err error, message ...interface{}) {
	if err != nil {
		log.Print(message, err)
	}

}

func Println(err error, message ...interface{}) {
	if err != nil {
		log.Println(message, err)
	}
}

func Fatal(err error, message ...interface{}) {
	if err != nil {
		log.Fatal(message, err)
	}

}

func Fatalln(err error, message ...interface{}) {
	if err != nil {
		log.Fatalln(message, err)
	}
}
