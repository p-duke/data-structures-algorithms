package main

import (
	"fmt"
	"sort"
)

// O(n^2) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	var result = []int{}
	for x := 0; x < len(array); x++ {
		for y := x + 1; y < len(array); y++ {
			sum := array[x] + array[y]
			if sum == target {
				result = append(result, array[x], array[y])
			}
		}
	}

	return result
}

// O(n) time | O(n) space
func TwoNumberSum(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		} else {
			nums[num] = true
		}
	}
	return []int{}
}

// O(nlog(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}
