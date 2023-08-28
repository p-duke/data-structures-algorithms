package main

import (
	"sort"
)

// Given an array of integers, nums, and an integer value, target,
// determine if there are any three integers in nums whose sum is equal to the target,
// that is, nums[i] + nums[j] + nums[k] == target.
// Return TRUE if three such integers exist in the array. Otherwise, return FALSE.
// Note: A valid triplet consists of elements with distinct indexes.

// Time Complexity 
// The time complexity of the function “findSumOfThree” is O(n^2) because of the nested for loop. 
// The outer loop runs n times, and the inner loop runs n times for each iteration of the outer loop. 
// The time complexity of the built-in Go function “sort.Ints” is O(n log n) because it uses quicksort algorithm. 
// The time complexity of the built-in Go functions used in the code, such as “len” and “sort”, is O(1) because they have constant time complexity.
// Therefore the total time complexity is O(nlog(n) + n^2) which is simplified to O(n^2)

// Space complexity is 0(1)
// The space complexity of the “findSumOfThree” function is O(1) because the function only uses a constant amount of extra space to store the “sum” variable
func FindSumOfThree(nums []int, target int) bool {
	var sum int

	if len(nums) < 3 {
		return false
	}

	// The nums need to be sorted for two pointer to work
	sort.Ints(nums)

  for idx := 0; idx < len(nums) - 1; idx++ {
    left := idx + 1;
    right := len(nums) - 1;

    for left < right {
      sum = nums[idx] + nums[left] + nums[right]

      if sum == target {
        return true
      } else if sum > target {
        right--
      } else if sum < target {
        left++
      }

    }
  }

	return false
}
