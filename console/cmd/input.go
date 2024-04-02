package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type any = interface{}

func Input(a ...any) string {
	fmt.Print(a...)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
