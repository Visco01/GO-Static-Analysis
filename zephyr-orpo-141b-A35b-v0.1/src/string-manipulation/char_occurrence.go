package main

import (
	"fmt"
)

func countChar(s string, c rune) int {
	count := 0
	for _, char := range s {
		if rune(char) == c {
			count++
		}
	}
	return count
}

func main() {
	s := "hello world"
	c := 'l'
	fmt.Println(countChar(s, c))
}
