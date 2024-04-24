package main

import (
	"fmt"
	"unicode"
)

func isAlphabetic(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlphabetic("helloworld"))
}
