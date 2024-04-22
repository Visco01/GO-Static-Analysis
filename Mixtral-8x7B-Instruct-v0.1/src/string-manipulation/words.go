package main

import (
	"fmt"
	"strings"
)

func countWords(str string) int {
	words := strings.Fields(str)
	return len(words)
}

func main() {
	str := "the house is white"
	fmt.Println(countWords(str))
}
