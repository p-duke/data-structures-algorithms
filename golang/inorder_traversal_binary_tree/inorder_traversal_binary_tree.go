package main

/*
Title: Inorder Traversal Binary Tree
Difficulty: Easy
Technique: ?
---

Given a binary tree, return the inorder traversal of its nodes' values.
Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
TODO: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
  result := make([]int, 0)
  traverse(root, &result)
  return result
}

// Recursive - Time complexity O(n) | Space complexity O(n)
func traverse(root *TreeNode, result *[]int) {
  if root != nil {
    traverse(root.Left, result)
    *result = append(*result, root.Val)
    traverse(root.Right, result)
  }
}

func main() {
  tests := []struct{
    root *TreeNode

  }{

  }
}
