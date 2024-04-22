package main

import (
	"fmt"
)

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < x {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low >= len(arr) || arr[low] != x {
		return -1
	}

	return low
}

func main() {
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 5))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 8))
}
