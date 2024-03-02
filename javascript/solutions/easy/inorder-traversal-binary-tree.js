/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var inorderTraversal = function(root) {
    const result = [];
    return traverse(root, result);
};

var traverse = function(node, result) {
    if (node === null) {
        return [];
    }

    traverse(node.left, result);
    result.push(node.val);
    traverse(node.right, result);
    return result;
};
