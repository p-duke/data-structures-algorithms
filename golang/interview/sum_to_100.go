/*

Question: Given an unordered list of integers between 1 and 99 (inclusive), write a function that generates all of the pairs of integers that sum to 100.

Examples:
sumTo100([89, 10]) == [[]]
sumTo100([1, 4, 31, 4, 2]) == [[]]
sumTo100([1, 99, 3, 2]) == [[1, 99]]
sumTo100([1, 99, 99, 1]) == [[1, 99], [1, 99]]
sumTo100([10, 20, 30, 40, 50, 60, 70, 80, 90]) == [[10, 90], [20, 80], [30, 70], [40, 60]]
sumTo100([50, 40, 95, 5, 60, 50]) == [[50, 50], [5, 95], [40, 60]]

// Problem
// input - unordered list of int between 1 - 99 (sort?) Yes
// target - 100 desired sum
// Multiple pairs equal the target

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(sumTo100([]int{10, 20, 30, 40, 50, 60, 70, 80, 90}))
	// fmt.Println(sumTo100([]int{50, 40, 95, 5, 60, 50}))
	fmt.Println(sumTo100([]int{50, 50, 50, 50, 50})) // [[50, 50], [50, 50]]
}


// Solution
// - First going to sort the list
// - Two pointer strategy - Left and Right pointer - O(n)
// - Left pointer - 0
// - Right pointer - len(arr) - 1
// - assign a var to sum left and right
// - Initiate a loop - left != right

// - check if sum == target then add left and right to result array
// 	- increment / decrement at the same time
// - if not check if sum > target then decrement right pointer
// - if not check if sum < target then increment left pointer

// O(n(log(n)) + n) time | O(n) space
// Reduced to O(n(log(n))) time
func sumTo100(input []int) [][]int {
	// Sort - O(n(log(n)))
	sort.Ints(input)
  target := 100
	var sum int
	result := make([][]int, 0)
	left, right := 0, len(input) - 1
	
	// O(n) time
	for left < right {
		sum = input[left] + input[right]
		if sum == target {
			temp := []int{input[left], input[right]}
			result = append(result, temp)
			left += 1
			right -=1
		}

		if sum > target {
			right -= 1
		}

		if sum < target {
			left += 1
		}
	}

	return result
}

