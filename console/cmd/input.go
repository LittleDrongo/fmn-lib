package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type any = interface{}

func Input(a ...any) string {
	fmt.Print(a...)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Password(a ...any) string {

	fmt.Print(a...)

	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	return string(bytePassword)
}
