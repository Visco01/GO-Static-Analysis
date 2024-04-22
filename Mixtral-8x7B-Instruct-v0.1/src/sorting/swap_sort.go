package main

import (
	"fmt"
)

func swapSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	sampleData := []int{5, 3, 1, 4, 2}
	fmt.Println("Original:", sampleData)
	swapSort(sampleData)
	fmt.Println("Sorted :", sampleData)
}
