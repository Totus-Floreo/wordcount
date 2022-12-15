package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	readedStr, err := readString()
	if err != nil && err != io.EOF {
		fmt.Println(fmt.Errorf("uncorrect string: %v", err))
		os.Exit(1)
	}
	fmt.Println(wordCount(readedStr))
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(str), nil
}

func wordCount(str string) int {
	if str == "\n" {
		return 0
	}
	words := strings.Split(str, " ")
	return len(words)
}
