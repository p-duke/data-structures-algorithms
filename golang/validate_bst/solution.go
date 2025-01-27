package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// Initial call with the entire range of possible integer values
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		// Base case: an empty node is valid
		return true
	}

	// Check if the current node value violates the BST property
	if node.Val <= min || node.Val >= max {
		return false
	}

	// Recursively check the left and right subtrees
	// Left subtree: max value is updated to current node's value
	// Right subtree: min value is updated to current node's value
	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
