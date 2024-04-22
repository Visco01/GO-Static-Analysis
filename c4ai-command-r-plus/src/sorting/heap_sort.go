package main

import (
	"fmt"
)

func heapify(arr []int, N int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < N && arr[left] > arr[largest] {
		largest = left
	}

	if right < N && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, N, largest)
	}
}

func heapSort(arr []int) {
	N := len(arr)

	for i := N / 2; i >= 0; i-- {
		heapify(arr, N, i)
	}

	for i := N - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	heapSort(arr)
	fmt.Println(arr)
}
