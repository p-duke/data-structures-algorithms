// Solution 1
function isMonotonic(array) {
	if (array.length === 0 || array.length == 1) return true;
	
	let monotonic = false;
	let last;
	let increase;
	let decrease;
	for (let i = 0; i < array.length; i++) {
		let curr = array[i];
		
		if (last && !decrease && last < curr) {
			increase = true;
			monotonic = true;
		} else if (last && !increase && curr < last) {
			decrease = true;
			monotonic = true;
		} else if (last === curr) {
			monotonic = true;
		} else if (increase && last > curr || decrease && curr > last) {
			return false;
		}
		
		last = curr;
	}
	
	return monotonic;
}

// Do not edit the line below.
exports.isMonotonic = isMonotonic;
