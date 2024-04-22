package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapify(arr []int) {
	h := &IntHeap(arr)
	heap.Init(h)
	for i := 0; i < cap(*h)-1; i++ {
		heap.Push(h, (*h)[len(*h)-1])
		(*h)[len(*h)-1] = (*h)[cap(*h)-1]
		(*h) = (*h)[:cap(*h)-1]
	}
}

func heapSort(arr []int) []int {
	heapify(arr)
	var h *IntHeap
	for h.Len() > 1 {
		h.Pop().(int)
		h.Swap(0, h.Len()-1)
		h.Pop()
		h.Push(&arr[h.Len()-1])
	}
	return arr
}

func main() {
	sampleData := []int{5, 3, 1, 4, 2}
	fmt.Println("Original:", sampleData)
	fmt.Println("Sorted :", heapSort(sampleData))
}
