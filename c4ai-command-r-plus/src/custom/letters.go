package main

import (
	"fmt"
	"unicode"
)

func isAlphabetic(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlphabetic("hello"))
}
