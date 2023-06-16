// Solution 1
function bubbleSort(array) {
  // Write your code here.
	// Loop over the array while until swaps is false
	// Once index === array.length - 1 then check swaps
  let swaps = false;
	let i = 0;
	while (i < array.length) {
		let curr = array[i];
		let next = array[i + 1];
		if (curr > next) {
			array[i] = next;
			array[i + 1] = curr;
			swaps = true;
		}
		
		i += 1;
		
		if (i === array.length - 1 && swaps) {
			swaps = false;
			i = 0;
		}
	}
	
	return array;
}

// Do not edit the line below.
exports.bubbleSort = bubbleSort;

// Solution 3 - Accepted
// Best: O(n) time | 0(1) space
// Average: O(n^2) time | O(1) space
// Worst: O(n^2) time | O(1) space
function bubbleSort(array) {
  // Write your code here.
	let isSorted = false;
	let counter = 0;
	while (!isSorted) {
		isSorted = true;
		for (let i = 0; i < array.length - 1 - counter; i++) {
			if (array[i] > array[i + 1]) {
				swap(i, i + 1, array);
				isSorted = false;
			}
		}
		counter += 1;
	}
	
	return array;
}

function swap(i, j, array) {
	const temp = array[j]
	array[j] = array[i];
	array[i] = temp;
}

// Do not edit the line below.
exports.bubbleSort = bubbleSort;

