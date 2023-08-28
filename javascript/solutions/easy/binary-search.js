// Solution 1
function binarySearch(array, target) {
  // Write your code here.
	// Problem
	// input is a sorted array of integers and target integer
	// If the target is in the array then return index
	// Else return -1
	
	// Solution
	// Recursively narrow down the array
	// Find the middle of the array
	// Compare to the target
	// If the element is larger than the target Then move left
	// If the element is less than the target Then move right
	// Make recursive call with new array
	let middle = Math.floor(array.length / 2);
	let el = array[middle];
	console.log(`el is ${el}`);
	if (el > target) {
		let newArr = array.slice(0, middle);
		console.log(`newArr is ${newArr}`);
		return binarySearch(newArr, target);
	} else if (el < target) {
		let newArr = array.slice(middle, array.length);
		console.log(`newArr is ${newArr}`);
		return binarySearch(newArr, target);
	} else if (el === target) {
		return middle;
	}
	
	return -1;
}

// Do not edit the line below.
exports.binarySearch = binarySearch;

// Accepted Solution - Iterative
// O(log(n)) time | O(1) space
function binarySearch(array, target) {
  // Write your code here.
	return search(array, target, 0, array.length -1);
}

function search(array, target, left, right) {
	while(left <= right) {
		const middle = Math.floor((left + right) / 2);
		const potentialMatch = array[middle];
		if (target === potentialMatch) {
			return middle;
		} else if (target < potentialMatch) {
			right = middle - 1;
		} else {
			left = middle + 1;
		}
	}
	return -1;
}

// Do not edit the line below.
exports.binarySearch = binarySearch;

// Accepted Solution - Recursive
// O(log(n)) time | O(log(n)) space
function binarySearch(array, target) {
  // Write your code here.
	return search(array, target, 0, array.length -1);
}

function search(array, target, left, right) {
	if (left > right) return -1;
	const middle = Math.floor((left + right) / 2);
	const potentialMatch = array[middle];
	if (target === potentialMatch) {
		return middle;
	} else if (target < potentialMatch) {
	  return search(array, target, left, middle - 1);
	} else {
	  return search(array, target, middle + 1, right);
	}
}

// Do not edit the line below.
exports.binarySearch = binarySearch;

