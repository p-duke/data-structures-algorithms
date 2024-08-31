package main

import "fmt"

func main() {
	h := &MinHeap{}

	h.Push(10)
	h.Push(20)
	h.Push(5)
	h.Push(15)
	h.Push(7)

	fmt.Println("Heap:", h.data)      // Heap: [5 7 10 20 15]
	fmt.Println("Min:", h.Pop())      // Min: 5
	fmt.Println("Heap after pop:", h.data) // Heap after pop: [7 15 10 20]
	fmt.Println("Min:", h.Pop())      // Min: 7
	fmt.Println("Heap after pop:", h.data) // Heap after pop: [10 15 20]
}
