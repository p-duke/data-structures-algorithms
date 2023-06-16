// Brute force to pass tests but not optimal or accepted solution
function findThreeLargestNumbers(array) {
	// Problem
	// input is an array of at least 3 integers
	// cannot sort array
	// returns a sorted array of 3 largets integers in array
	// If the top 3 numbers are duplicates they should be included
	
	// Solution
	// Loop over entire array - O(n) time
	// create first, second, third variables = undefined
	// if curr > first then first = curr
	// if curr > second then second = curr
	// if curr > third then third = curr
	let first = array[0];
	let second = array[1];
	let third = array[2];
	
	if (array.length === 3) {
		return [second, third, first];
	}
	
	for (let i = 3; i < array.length; i++) {
		const curr = array[i];
		if (curr > first) {
			if (first > second) {
				if (second >= third) {
					third = second;
				}
				second = first;
			}
			first = curr;
		} else if (curr > second) {
			if (second > third) {
				third = second;
			}
			second = curr;
		} else if (curr > third) {
			third = curr;
		}
	}
	
	return [third, second, first];
}

// Do not edit the line below.
exports.findThreeLargestNumbers = findThreeLargestNumbers;

// Solution 2 - Not finished
function findThreeLargestNumbers(array) {
	/* Solution
	 - Loop through the entire array
	 - Declare an array to hold the three largest numbers
	 - In each iteration check if the number is larger than
	 final[0], final[1], final[2] from smallest to largest
	 - If curr > final[2] then move 2 to 1
	 - If curr > final [1] then move 1 to 0
	 - If curr > final [0] then remove 0 and replace with 1
	*/
	const final = [undefined, undefined, undefined];
	for (let i = 0; i < array.length; i++) {
		const curr = array[i];
		if (!final[2] || curr > final[2]) {
			if (!final[2]) {
				final[2] = curr;
				continue;
			}

			final[1] = final[2];
			final[2] = curr;
		}

		if (!final[1] || curr > final[1]) {
			if (!final[1]) {
				final[1] = curr;
				continue;
			}

			final[0] = final[1];
			final[1] = curr;
		}

		if (curr > final[0]) {
			if (!final[0]) {
				final[0] = curr;
				continue;
			}

			final[0] = curr;
		}
	}

	return final;
}

// Do not edit the line below.
exports.findThreeLargestNumbers = findThreeLargestNumbers;

// Accepted Solution
// O(n) time | O(1) space
function findThreeLargestNumbers(array) {
	const threeLargest = [null, null, null];
	for (let i = 0; i < array.length; i++) {
		updateLargest(threeLargest, array[i]);
	}
	
	return threeLargest;
}

function updateLargest(threeLargest, num) {
	if (threeLargest[2] === null || num > threeLargest[2]) {
		shiftAndUpdate(threeLargest, num, 2);
	} else if (threeLargest[1] === null || num > threeLargest[1]) {
		shiftAndUpdate(threeLargest, num, 1);
	} else if (threeLargest[0] === null || num > threeLargest[0]) {
		shiftAndUpdate(threeLargest, num, 0);
	}
}

function shiftAndUpdate(array, num, idx) {
	for (let i = 0; i < idx + 1; i++) {
		if (i === idx) {
			array[i] = num;
		} else {
			array[i] = array[i + 1];
		}
	}
}

// Do not edit the line below.
exports.findThreeLargestNumbers = findThreeLargestNumbers;

