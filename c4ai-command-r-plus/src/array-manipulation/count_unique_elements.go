package main

import (
	"fmt"
)

func countUniqueElements(arr []int) int {
	uniqueMap := make(map[int]bool)
	for _, num := range arr {
		uniqueMap[num] = true
	}
	return len(uniqueMap)
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(countUniqueElements(arr))
}
