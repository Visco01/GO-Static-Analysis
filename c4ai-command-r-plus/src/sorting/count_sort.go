package main

import (
	"fmt"
)

func countSort(arr []int) []int {
	maxVal := 0
	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
	}

	count := make([]int, maxVal+1)
	for _, num := range arr {
		count[num]++
	}

	sortedIndex := 0
	for i, freq := range count {
		for j := 0; j < freq; j++ {
			arr[sortedIndex] = i
			sortedIndex++
		}
	}

	return arr
}

func main() {
	arr := []int{2, 5, 3, 7, 2, 3, 6}
	fmt.Println(countSort(arr))
}
