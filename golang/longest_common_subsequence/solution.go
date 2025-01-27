package main

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	// Step 1: Create a DP table with dimensions (len(text1)+1) x (len(text2)+1)
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Step 2: Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If characters match, add 1 to the LCS length of the previous state
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// If characters don't match, take the maximum LCS length from excluding one character
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Step 3: The answer is in dp[m][n], which represents the LCS of the two full strings
	return dp[m][n]
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage:
	text1 := "abcde"
	text2 := "ace"
	result := longestCommonSubsequence(text1, text2)
	fmt.Printf("The length of the Longest Common Subsequence is: %d\n", result)
}

