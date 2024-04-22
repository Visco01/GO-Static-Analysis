package main

import (
	"fmt"
)

func countingUniqueElements(arr []int) int {
	m := map[int]int{}
	count := 0

	for i := range arr {
		m[arr[i]]++
	}

	for _, v := range m {
		if v == 1 {
			count += 1
		}
	}

	return count
}

func main() {
	fmt.Println(countingUniqueElements([]int{1, 2, 3, 4, 4, 4, 5}))
	fmt.Println(countingUniqueElements([]int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 6}))
}
