package cmd

import (
	"fmt"
)

func Input(str string) string {

	var result string
	fmt.Print(str)
	fmt.Scanf("%s\n", &result)
	return result
}
