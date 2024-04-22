package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func main() {
	unsortedArr := []int{5, 3, 1, 4, 2}
	insertionSort(unsortedArr)
	fmt.Println(unsortedArr)
}
