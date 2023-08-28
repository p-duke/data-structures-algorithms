function nodeDepths(root) {
	const depths = inOrder(root, 0, {});
	let sum = 0;
	for (const depth in depths) {
		sum += depths[depth];
	}
	return sum;
}

const inOrder = (node, depth, depths) => {
	depths[node.value] = depth;
	node.left && inOrder(node.left, depth += 1, depths);
  // We only want to count one per depth - so if not left then right
	if (node.left === null) {
		depth += 1;
	}
	node.right && inOrder(node.right, depth, depths); 
	return depths;
}

// This is the class of the input binary tree.
class BinaryTree {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

// Do not edit the line below.
exports.nodeDepths = nodeDepths;

