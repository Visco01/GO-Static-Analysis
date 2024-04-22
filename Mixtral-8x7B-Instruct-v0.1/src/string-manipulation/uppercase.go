package main

import (
	"fmt"
	"strings"
)

func makeUppercase(str string, length int) string {
	upper := make([]byte, length)
	for i := 0; i < length; i++ {
		upper[i] = byte(strings.ToUpper(string(str[i]))[0])
	}
	return string(upper)
}

func main() {
	str := "ciao"
	fmt.Println(makeUppercase(str, len(str)))
}
