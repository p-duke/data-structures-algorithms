package main

// Heap Array Operations
// root - 0
// current - i
// parent - (i -1)/2
// left child - (2*i) + 1
// right child - (2*i) + 2
// the last non-leaf - (len(array) - 2) / 2

type MinHeapSolution struct {
	data []int
}

// Push method adds data to the underlying array
// and bubbles up the value
func (h *MinHeapSolution) Push(x int) {
	h.data = append(h.data, x)
	h.bubbleUp(len(h.data)-1)
}

// BubbleUp
func (h *MinHeapSolution) bubbleUp(i int) {
	for i > 0 && h.data[i] < h.data[h.parent(i)] {
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		i = h.parent(i)
	}
}

// Pop
func (h *MinHeapSolution) Pop() int {
	if len(h.data) == 0 {
		return -1
	}

	root := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.bubbleDown(0)
	return root
}

// bubbleDown
func (h *MinHeapSolution) bubbleDown(i int) {
	left := h.leftChild(i)
	right := h.rightChild(i)
	smallest := i

	if left < len(h.data) && h.data[left] < h.data[smallest] {
		smallest = left
	}

	if right < len(h.data) && h.data[right] < h.data[smallest] {
		smallest = right
	}

	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.bubbleDown(smallest)
	}
}

func (h *MinHeapSolution) Peek() int {
	if len(h.data) == 0 {
		return -1
	}

	return h.data[0]
}

// Helper Methods
func (h *MinHeapSolution) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeapSolution) leftChild(i int) int {
	return (2 * i) + 1
}

func (h *MinHeapSolution) rightChild(i int) int {
	return (2 * i) + 2
}

