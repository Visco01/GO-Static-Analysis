package main

import (
	"fmt"
)

func occurrenceCount(arr []int, target int) int {
	count := 0
	for _, num := range arr {
		if num == target {
			count++
		}
	}
	return count
}

func main() {
	arr := []int{1, 2, 2, 3, 2, 4, 2, 5}
	target := 2
	fmt.Println(occurrenceCount(arr, target))
}
