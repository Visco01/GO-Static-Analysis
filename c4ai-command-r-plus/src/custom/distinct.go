package main

import (
	"fmt"
)

func areDistinct(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println(areDistinct(arr1))
}
