package main

import (
  "fmt"
  "reflect"
)

/* 
Title: Validate Subsequence
Difficulty: Easy
Pattern: N/A
---

Question: Given two non-empty arrays of integers, write a function that determines whether the second array is a subsequence of the first one.

A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array but that are in the same order as they appear in the array.
For instance, the numbers [1, 3, 4] form a subsequence of the array [1, 2, 3, 4], and so do the numbers [2, 4].
Note that a single number in an array and the array itself are both valid subsequences of the array.

Sample Input
array = [5, 1, 22, 25, 6, -1, 8, 10]
sequence = [1, 6, -1, 10]

Sample Output
true

---

Hint: Hints
Hint 1

You can solve this question by iterating through the main input array once.

Hint 2

Iterate through the main array, and look for the first integer in the potential subsequence. If you find that integer, keep on iterating through the main array, but now look for the second integer in the potential subsequence. Continue this process until you either find every integer in the potential subsequence or you reach the end of the main array.

Hint 3

To actually implement what Hint #2 describes, you'll have to declare a variable holding your position in the potential subsequence. At first, this position will be the 0th index in the sequence; as you find the sequence's integers in the main array, you'll increment the position variable until you reach the end of the sequence.

Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the array

*/

func isValidSubsequence(array []int, sequence []int) bool {
  // Problem
  // Check if subsequence is not adjacent but in the array
  // Single number in an array and the array itself are valid subsequences

  // Solution
  // Loop over every element of the array
  // Keep track of a sequence ID
  // If the sequence index el matches the array el then increment to next sequence ID
  // Else next array element
  // If sequence length equals seqId we can exit the loop
  // Return whether len of sequence equals seqId
  // Time complexity O(n) | Space complexity O(1)
  seqId := 0
  for _, v := range array {
    if seqId == len(sequence) {
      break
    }

    if v == sequence[seqId] {
      seqId++
    }
  }

  return seqId == len(sequence)
}

func main() {
  tests := []struct{
    array []int
    sequence []int
    want bool
  }{
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, -1, 10}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 25, 6, -1, 8, 10}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 6, -1, 8, 10}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{22, 25, 6}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, 10}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 10}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, -1, 8, 10}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{25}, want: true },
    { array: []int{1, 1, 1, 1, 1}, sequence: []int{1, 1, 1}, want: true },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 25, 6, -1, 8, 10, 12}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{4, 5, 1, 22, 25, 6, -1, 8, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 23, 6, -1, 8, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 22, 25, 6, -1, 8, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 22, 6, -1, 8, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, -1, -1}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, -1, -1, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, -1, -2}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{26}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 25, 22, 6, -1, 8, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 26, 22, 8}, want: false },
    { array: []int{1, 1, 6, 1}, sequence: []int{1, 1, 1, 6}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, -1, 10, 11, 11, 11, 11}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{5, 1, 22, 25, 6, -1, 8, 10, 10}, want: false },
    { array: []int{5, 1, 22, 25, 6, -1, 8, 10}, sequence: []int{1, 6, -1, 5}, want: false },
  }

  for _, tc := range tests {
    got := isValidSubsequence(tc.array, tc.sequence)
    if !reflect.DeepEqual(tc.want, got) {
      fmt.Printf("FAIL! Expected: %v --- Got: %v\n", tc.want, got)
    } else {
      // fmt.Printf("PASS Expected: %v --- Got: %v\n", tc.want, got)
    }
  }
}

