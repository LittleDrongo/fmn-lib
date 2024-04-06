package cmd

import (
	"fmt"
	"time"
)

func Timeout(loading []string, second int, message ...any) {

	timeout := time.Duration(second) * time.Second

	start := time.Now()
	for {
		elapsed := time.Since(start)
		if elapsed >= timeout {
			fmt.Println()
			return
		}
		for _, frame := range loading {
			fmt.Printf("\r%s %s", message, frame)
			time.Sleep(timeout / time.Duration(len(loading)))
		}
	}
}

func TimeoutRun(loading []string, second int, code func(), message ...any) {

	timeout := time.Duration(second) * time.Second

	start := time.Now()
	for {
		elapsed := time.Since(start)
		if elapsed >= timeout {
			fmt.Println()
			code()
			return
		}
		for _, frame := range loading {
			fmt.Printf("\r%s %s", message, frame)
			time.Sleep(timeout / time.Duration(len(loading)))
		}
	}
}
