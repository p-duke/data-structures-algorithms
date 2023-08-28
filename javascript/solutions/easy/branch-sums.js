// This is the class of the input root.
// Do not edit it.
class BinaryTree {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

// Solution 1
function calculateSums(node, runningSum, sums) { 
	if (!node) return;
	let newRunningSum = runningSum + node.value;
	if (!node.left && !node.right) {
		sums.push(newRunningSum);
	}
	
	calculateSums(node.left, newRunningSum, sums);
	calculateSums(node.right, newRunningSum, sums);
} 

function branchSums(root) {
	let sums = [];
	calculateSums(root, 0, sums);
	return sums;
}

// Do not edit the lines below.
exports.BinaryTree = BinaryTree;
exports.branchSums = branchSums;

