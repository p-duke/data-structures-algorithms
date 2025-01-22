package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Create a dp array where dp[i] represents the LIS ending at index i
	dp := make([]int, n)

	// Initialize dp array to 1, as every element is an LIS of length 1 by itself
	for i := range dp {
		dp[i] = 1
	}

	// Fill dp array based on LIS conditions
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// Find the maximum value in dp array, which is the LIS
	maxLength := 0
	for _, length := range dp {
		maxLength = max(maxLength, length)
	}

	return maxLength
}

// Utility function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Length of LIS:", lengthOfLIS(nums))
}

