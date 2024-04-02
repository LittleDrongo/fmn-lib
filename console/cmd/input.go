package cmd

import (
	"fmt"
)

type any = interface{}

func Input(a any) string {

	var result string
	fmt.Print(a)
	fmt.Scanf("%s\n", &result)
	return result
}
