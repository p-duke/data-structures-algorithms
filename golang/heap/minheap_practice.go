package main

import "container/heap"

type MinHeap []int

// heap.Interface methods Push, Pop
func (mh *MinHeap) Push(num any) {
	*mh = append(*mh, num.(int))
}

func (mh *MinHeap) Pop() any {
	oldMh := *mh
	oldLen := len(*mh)
	popped := oldMh[oldLen-1]
	*mh = oldMh[:oldLen-1]

	return popped
}

// heap.Sort methods Len, Less, Swap
func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func buildHeapByInit(nums []int) *MinHeap {
	mh := &MinHeap{}
	*mh = nums
	heap.Init(mh)
	return mh
}

func buildHeapByPush(nums []int) *MinHeap {
	mh := &MinHeap{}
	for _, num := range nums {
		heap.Push(mh, num)
	}
	return mh
}

