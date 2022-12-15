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
	words := os.Args[1:]
	if strings.TrimSpace(words[0]) == "" {
		return 0
	}
	return len(words)
}
