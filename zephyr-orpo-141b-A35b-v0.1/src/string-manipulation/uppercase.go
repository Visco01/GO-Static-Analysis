package main

import (
	"fmt"
)

func toUpperCase(s string) string {
	return fmt.Sprintf("%s", []rune(s))
}

func main() {
	s := "ciao"
	fmt.Println(toUpperCase(s))
}
