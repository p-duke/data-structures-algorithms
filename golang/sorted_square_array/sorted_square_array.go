package main

import (
	"fmt"
	"math"
	"reflect"
)

/*

Title: Sorted Squared Array
Difficulty: Easy
Technique: Two pointers

---
Question: Write a function that takes in a non-empty array of integers that are sorted in ascending order and returns a new array of the same length with the squares of the original integers also sorted in ascending order.

Sample Input
array = [1, 2, 3, 5, 6, 8, 9]

Sample Output
[1, 4, 9, 25, 36, 64, 81]
---

Hint: Hints
Hint 1

While the integers in the input array are sorted in increasing order, their squares won't necessarily be as well, because of the possible presence of negative numbers.

Hint 2

Traverse the array value by value, square each value, and insert the squares into an output array. Then, sort the output array before returning it. Is this the optimal solution?

Hint 3

To reduce the time complexity of the algorithm mentioned in Hint #2, you need to avoid sorting the ouput array. To do this, as you square the values of the input array, try to directly insert them into their correct position in the output array.

Hint 4

Use two pointers to keep track of the smallest and largest values in the input array. Compare the absolute values of these smallest and largest values, square the larger absolute value, and place the square at the end of the output array, filling it up from right to left. Move the pointers accordingly, and repeat this process until the output array is filled.

Optimal Space & Time Complexity
O(n) time | O(n) space - where n is the length of the input array
*/

// Question: Write a function that takes in a non-empty array of integers that are sorted in ascending order and returns a new array of the same length with the squares of the original integers also sorted in ascending order.
func sortedSquaredArray(arr []int) []int {
	// Problem
	// Input - non-empty array of int (sorted asc)
	// Return - new array of same length with squares of original integers (sorted asc)
	// Edge case - negative numbers - [-2, -1]

	// Solution - two pointer
	// assign left and right pointer to first and last element
	// loop while left != right
	// Take absolute of left and right and compare
	// Take bigger of two, square, append and increment/decrement one pointer
	var result []int
	left := 0
	right := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		leftAbs := math.Abs(float64(arr[left]))
		rightAbs := math.Abs(float64(arr[right]))

		if leftAbs > rightAbs {
			leftSquared := leftAbs * leftAbs
			result = append([]int{int(leftSquared)}, result...)
			left++
		} else {
			rightSquared := rightAbs * rightAbs
			result = append([]int{int(rightSquared)}, result...)
			right--
		}
	}

	return result
}

func main() {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, 3, 5, 6, 8, 9}, want: []int{1, 4, 9, 25, 36, 64, 81}},
		{input: []int{1}, want: []int{1}},
		{input: []int{1, 2}, want: []int{1, 4}},
		{input: []int{1, 2, 3, 4, 5}, want: []int{1, 4, 9, 16, 25}},
		{input: []int{0}, want: []int{0}},
		{input: []int{10}, want: []int{100}},
		{input: []int{-1}, want: []int{1}},
		{input: []int{-2, -1}, want: []int{1, 4}},
		{input: []int{-10}, want: []int{100}},
		{input: []int{-10, -5, 0, 5, 10}, want: []int{0, 25, 25, 100, 100}},
		{input: []int{-7, -3, 1, 9, 22, 30}, want: []int{1, 9, 49, 81, 484, 900}},
		{input: []int{-50, -13, -2, -1, 0, 0, 1, 1, 2, 3, 19, 20}, want: []int{0, 0, 1, 1, 1, 4, 4, 9, 169, 361, 400, 2500}},
		{input: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, want: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{input: []int{-1, -1, 2, 3, 3, 3, 4}, want: []int{1, 1, 4, 9, 9, 9, 16}},
		{input: []int{-3, -2, -1}, want: []int{1, 4, 9}},
	}

	for _, tc := range tests {
		got := sortedSquaredArray(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("Failed! Got: %v, Want: %v\n", got, tc.want)
		} else {
			fmt.Printf("PASSED! Got: %v, Want: %v\n", got, tc.want)
		}
	}
}
