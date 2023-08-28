class BinaryTree {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

/**
* @param {TreeNode} root
*/
const inOrder = (root) => {
  const nodes = []
  if (root) {
    inOrder(root.left)
    nodes.push(root.val)
    inOrder(root.right)
  }
  return nodes
}
// for our example tree, this returns [1,2,3,4,5,6]

/**
* @param {TreeNode} root
*/
const postOrder = (root) => {
  const nodes = []
  if (root) {
    postOrder(root.left)
    postOrder(root.right)
    nodes.push(root.val)
  }
  return nodes
}
// for our example tree, this returns [1,3,2,6,5,4]

/**
* @param {TreeNode} root
*/
const preOrder = (root) => {
  const nodes = []
  if (root) {
    nodes.push(root.val)
    preOrder(root.left)
    preOrder(root.right)
  }
  return nodes
}
// for our example tree, this returns [4,2,1,3,5,6]
