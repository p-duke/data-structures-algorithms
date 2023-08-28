function moveElementToEnd(array, toMove) {
  // Problem
	// array of integers and integer
	// Move all instances of toMove to the end of the array
	// mutate the input array
	
	// Solution
	// Loop over array
	// arr.splice(arr.indexOf(toMove), 1)
	// arr.push(toMove);
	for (let i = 0; i < array.length; i++) {
		const el = array[i];
		if (el === toMove) {
			array.splice(array.indexOf(toMove), 1);
			array.push(toMove);
		}
	}
	
	return array;
}

// Do not edit the line below.
exports.moveElementToEnd = moveElementToEnd;

