// Solution 1
function findClosestValueInBst(tree, target) {
	let curr = tree;
	let diff = Math.abs(target - curr.value);
	
	if (diff === 0) return curr.value;
	
	let lowestDiff = diff;
	let lowest = curr.value;
	
	while (true) {
		let left = curr.left ? curr.left : null;
		let right = curr.right ? curr.right : null;
		let diff = Math.abs(target - curr.value);
		
		if (diff < lowestDiff) {
			lowest = curr.value;
			lowestDiff = diff;
		}	
		
		if (right && curr.value < target) {
			curr = curr.right
		} else if (left && curr.value > target) {
			curr = curr.left;
		} else {
			return lowest;
		}
	}
}

// This is the class of the input tree. Do not edit.
class BST {
  constructor(value) {
    this.value = value;
    this.left = null;
    this.right = null;
  }
}

// Do not edit the line below.
exports.findClosestValueInBst = findClosestValueInBst;

