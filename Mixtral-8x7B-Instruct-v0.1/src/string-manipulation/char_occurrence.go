package main

import (
	"fmt"
)

func countCharOccurrences(str string, char rune) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == char {
			count++
		}
	}
	return count
}

func main() {
	str := "hello world"
	char := 'l'
	fmt.Println(countCharOccurrences(str, char))
}
