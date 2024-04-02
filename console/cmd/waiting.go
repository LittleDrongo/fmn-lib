package cmd

import (
	"fmt"
	"time"
)

func Waiting(animation []string, message ...interface{}) {
	for {
		for _, frame := range animation {
			fmt.Printf("\r%v %s", message, frame)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
