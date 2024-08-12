package main

import (
	"fmt"
	"reflect"
)

func intersection(nums1 []int, nums2 []int) []int {
	return []int{}
}

func main() {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{nums1: []int{}, nums2: []int{}, want: []int{}},
	}

	for _, tc := range tests {
		got := intersection(tc.nums1, tc.nums2)

		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAILED! got: %v want: %v", got, tc.want)
		} else {
			fmt.Printf("PASSED! got: %v want: %v", got, tc.want)
		}
	}
}
