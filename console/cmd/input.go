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

func Rune(a ...any) rune {
	fmt.Print(a...)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return 0
	}
	return char
}

func PressKey(a ...any) rune {

	fmt.Print(a...)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	char, err := readSingleChar()
	if err != nil {
		return 0
	}

	return char
}

func readSingleChar() (rune, error) {
	var buf [1]byte
	_, err := os.Stdin.Read(buf[:])
	if err != nil {
		return 0, err
	}
	return rune(buf[0]), nil
}
