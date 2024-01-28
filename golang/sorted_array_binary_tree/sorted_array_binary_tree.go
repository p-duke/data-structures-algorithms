package main

import (
  "fmt"
  "reflect"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
 }

// Time complexity O(n) | Space complexity O(n)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

// NOT FINISHED
func main() {
  tests := []struct{
    nums []int
    want []int
  }{
    { nums: []int{-10,-3,0,5,9}, want: []int{-10,-3,0,5,9} }
  }

  for _, tc := range tests {
    if !reflect.DeepEqual(got, want) {
      fmt.Println("FAILED! got: %d, want: %d", got, want)
    } else {
      fmt.Println("PASSED! got: %d, want: %d", got, want)
    }
  }
}
