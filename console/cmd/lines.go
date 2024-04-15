package cmd

import "fmt"

func ClearPreviousLine() {
	fmt.Print("\033[F\033[K")
}

func ClearPreviousLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[F\033[K")
	}
}
