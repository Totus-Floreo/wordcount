package main

import (
	"fmt"
	"os"
	"strings"
)

// start : ./wordcount 'go is awesome'

func main() {
	fmt.Println(wordsCount())
}

func wordsCount() int {
	enter := os.Args[1:]
	words := strings.Split(enter[0], " ")
	if strings.TrimSpace(words[0]) == "" {
		return 0
	}
	return len(words)
}
