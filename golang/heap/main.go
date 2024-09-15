package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := buildHeapByInit(nums)

	fmt.Printf("Heap: %v should be [1 10 9 22 31 15 40 25 91]\n", minHeap)
	fmt.Printf("Min: %d should be 1\n", (*minHeap)[0])

	elementToRemove := heap.Pop(minHeap)
	fmt.Println("removed element by Pop method: ", elementToRemove)

	fmt.Printf("Min: %d should be 9\n", (*minHeap)[0])

	fmt.Printf("Current minHeap %v should be [9 10 15 22 31 91 40 25]\n", minHeap)
	(*minHeap)[2] = 51
	fmt.Println("Modify min heap value at index 2: ", *minHeap)

	heap.Fix(minHeap, 2)
	fmt.Println("Fix min heap: ", *minHeap)

	heap.Push(minHeap, 3)
	fmt.Println("Pushing 3 onto minHeap", *minHeap)

	fmt.Printf("Min: %d should be 3\n", (*minHeap)[0])
}
