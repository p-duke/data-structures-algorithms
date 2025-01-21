package main

import (
	"fmt"
)

func maxArea(height []int) int {
	// Initialize two pointers: left at the beginning, right at the end of the array
	left, right := 0, len(height)-1
	maxArea := 0 // Variable to store the maximum area found

	// Loop while the two pointers don't overlap
	for left < right {
		// Calculate the current area
		// Height is determined by the shorter of the two lines
		h := min(height[left], height[right])
		// Width is the distance between the two pointers
		width := right - left
		// Calculate area
		currentArea := h * width

		// Update maxArea if the current area is larger
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// Move the pointer pointing to the shorter line inward
		// This is because moving the taller line inward will not help increase the area
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	// Return the maximum area found
	return maxArea
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example input
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// Call the maxArea function and print the result
	fmt.Println("Maximum area:", maxArea(height))
}

