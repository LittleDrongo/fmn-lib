package cmd

import (
	"fmt"
	"time"
)

func Timeout(animation []string, timeout time.Duration, a ...any) {
	start := time.Now()
	for {
		elapsed := time.Since(start)
		if elapsed >= timeout {
			fmt.Println()
			return
		}
		for _, frame := range animation {
			fmt.Printf("\r%s %s", a, frame)
			time.Sleep(timeout / time.Duration(len(animation)))
		}
	}
}
