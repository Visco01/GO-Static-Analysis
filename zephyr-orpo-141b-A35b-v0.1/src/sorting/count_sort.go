package main

import (
	"fmt"
)

func countSort(arr []int) {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	count := make([]int, max+1)
	for _, num := range arr {
		count[num]++
	}

	index := 0
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			arr[index] = i
			index++
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	countSort(arr)
	fmt.Println("Sorted array:", arr)
}
