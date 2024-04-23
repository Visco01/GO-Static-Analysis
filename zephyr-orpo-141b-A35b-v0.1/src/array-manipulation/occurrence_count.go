package main

import (
	"fmt"
)

func countOccurrences(arr []int, target int) int {
	count := 0
	for _, num := range arr {
		if num == target {
			count++
		}
	}
	return count
}

func main() {
	arr := []int{1, 2, 3, 2, 4, 2, 5}
	target := 2
	count := countOccurrences(arr, target)
	fmt.Printf("The number %d appears %d times in the array.\n", target, count)
}
