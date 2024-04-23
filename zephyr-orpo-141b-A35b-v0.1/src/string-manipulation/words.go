package main

import (
	"fmt"
	"strings"
)

func countWords(s string) int {
	return len(strings.Fields(s))
}

func main() {
	s := "the house is white"
	fmt.Println(countWords(s))
}
