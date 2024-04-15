package cmd

import "fmt"

func ClearPreviousLine() {
	fmt.Println("\033[F\033[K")
}

func ClearPreviousLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Println("\033[F\033[K")
	}
}
