/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    return dfs(root, map[int]bool{}, k)
}

func dfs(node *TreeNode, tracker map[int]bool, target int) bool {
    if node == nil {
        return false
    }

    diff := target - node.Val

    if _, ok := tracker[diff]; ok {
        return true
    }

    tracker[node.Val] = true

    left := dfs(node.Left, tracker, target)
    right := dfs(node.Right, tracker, target)

    return left || right
}
