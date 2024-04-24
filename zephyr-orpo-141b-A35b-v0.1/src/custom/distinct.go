package main

import (
	"fmt"
)

func isDistinct(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(isDistinct(arr))
}
