/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := height(root.Left)
	right := height(root.Right)

	if int(math.Abs(float64(left-right))) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}
