package cmd

import (
	"fmt"
	"time"
)

func Timeout(loading []string, timeout time.Duration, message ...any) {
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

func TimeoutRun(loading []string, timeout time.Duration, code func(), message ...any) {
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
