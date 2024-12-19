package main

import (
	"fmt"
	"reflect"
)

// INCOMPLETE - SKIPPED
func circularArrayLoop(nums []int) bool {
	// On each loop
	// - Move the slow pointer 1 forward/backward
	// - Move the fast pointer 2 forward/backward
	// - If the loop changes direction move onto the next number
	// - Check if the pointers meet at each iteration - if yes return true
	for i := 0; i < len(nums)-1; i++ {
		var curr int = nums[i]
		var slowIdx int = i + curr
		var fastIdx int = nums[slowIdx]

		fmt.Println("slow", slowIdx)
		fmt.Println("fast", fastIdx)
	}

	return false
}

func main() {
	tests := []struct {
		arr  []int
		want bool
	}{
		// {arr: []int{1, 4, 3, 2, 1}, want: true},
		// {arr: []int{3,-3,2,-2}, want: false},
		// {arr: []int{-2, -3, 1, -3, 2}, want: true},
		// {arr: []int{5, -1, 1, 1, -7, -9}, want: false},
		// {arr: []int{2, 5, -4, 3, -1, 4}, want: false},

		{arr: []int{1, 3, -2, -4, 1}, want: true}, // Cycle starts w/ 1
		// {arr: []int{2,1,-1,-2}, want: false},
		// {arr: []int{5,4,-2,-1,3}, want: true},
		// {arr: []int{1,2,-3,3,4,7,1}, want: true},
		// {arr: []int{3,3,1,-1,2}, want: true},
	}

	for _, tc := range tests {
		got := circularArrayLoop(tc.arr)
		if !reflect.DeepEqual(tc.want, got) {
			fmt.Printf("FAIL! Assert: %d, Expected: %t --- Got: %t\n", tc.arr, tc.want, got)
		} else {
			fmt.Printf("PASS Assert: %d, Expected: %t --- Got: %t\n", tc.arr, tc.want, got)
		}
	}
}
