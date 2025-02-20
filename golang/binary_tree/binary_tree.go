package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) levelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root} // Initialize queue with root node
	result := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue) // Number of nodes at the current level
		level := []int{}        // Slice to store values of nodes at this level

		for i := 0; i < levelSize; i++ {
			// Dequeue the first node
			current := queue[0]
			queue = queue[1:]

			// Add node value to current level
			level = append(level, current.Val)

			// Enqueue left and right children if they exist
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		// Add the current level to the result
		result = append(result, level)
	}

	return result
}

func main() {

}
