package main

import (
	"fmt"
	"sort"
)

func countUniqueElements(arr []int) int {
	sort.Ints(arr)
	uniqueCount := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			uniqueCount++
		}
	}
	return uniqueCount
}

func main() {
	arr := []int{1, 2, 3, 2, 4, 2, 5}
	uniqueCount := countUniqueElements(arr)
	fmt.Printf("There are %d unique elements in the array.\n", uniqueCount)
}
