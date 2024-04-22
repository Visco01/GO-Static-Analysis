package main

import (
	"fmt"
)

func occurrenceCount(arr []int, target int) int {
	count := 0
	for _, value := range arr {
		if value == target {
			count++
		}
	}
	return count
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	target := 5
	fmt.Printf("Number of occurrences of %d in arr: %d\n", target, occurrenceCount(arr, target))
}
