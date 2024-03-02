/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isBalanced = function (root) {
    if (root === null) {
        return true;
    }

    return height(root).balanced;
};

var height = function (node) {
    if (!node) {
        return { height: 0, balanced: true };
    }

    const leftResult = height(node.left);
    const rightResult = height(node.right);

    const currentHeight = 1 + Math.max(leftResult.height, rightResult.height);
    const currentBalanced = Math.abs(leftResult.height - rightResult.height) <= 1 &&
        leftResult.balanced &&
        rightResult.balanced;

    return { height: currentHeight, balanced: currentBalanced };
}

