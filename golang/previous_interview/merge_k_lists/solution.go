package main

import (
	"fmt"
	"reflect"
)

// Solution
// - declare a result []int
// - loop over the slices
// - work with i and j
// - on each loop compare i < j and push into slice and increment
// - else push j into slice and increment j
// - potential leftover elements and merge into result
// O(n x m) time | O(n) space (merge sort)
// n = total nubmer of arrays
// m = total number of elements with those arrays
// O(n log(n)) - naive sort
func iterate(arrs [][]int) []int {

	result := arrs[0]
	for i := 1; i < len(arrs); i++ {
		result = merge(result, arrs[i])
	}
	return result
}

// Example
// {3, 9},
// {0, 5, 8}, and
// {2, 6, 7, 8}.
//
// The output of the function that provides a valid solution should be
// {0, 2, 3, 5, 6, 7, 8, 8, 9}.

// O(n) time | O(n) space
func merge(nums1, nums2 []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}
	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}
	return result
}
func main() {
	arrs := [][]int{
		{3, 9},
		{0, 5, 8},
		{2, 6, 7, 8},
	}
	want := []int{0, 2, 3, 5, 6, 7, 8, 8, 9}
	got := iterate(arrs)
	if !reflect.DeepEqual(got, want) {

		fmt.Printf("FAILED! got %v, want %v", got, want)
	} else {
		fmt.Printf("PASSED!")
	}
}
