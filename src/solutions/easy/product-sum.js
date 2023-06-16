// Tip: You can use the Array.isArray function to check whether an item
// is a list or an integer.

// Solution 1 - INCOMPLETE - Ran out of time
function productSum(array, total = 0, depth = 1) {
	if (!array.length) return total;
	
	for (let i = 0; i < array.length; i++) {
		let el = array[i];
		if (typeof el === 'number') {
			total += el;
		}
		
		if (Array.isArray(el)) {
			depth += 1;
			total += depth;
			let subSum = productSum(el, 0, depth);
			total *= subSum;
		}
	}
	
	return total;
}

// Do not edit the line below.
exports.productSum = productSum;

// Accepted Solution
function productSum(array, multiplier = 1) {
  let sum = 0;
  for (let i = 0; i < array.length; i++) {
    let el = array[i];
    if (Array.isArray(el)) {
      sum += productSum(el, multiplier + 1);
    } else {
      sum += el;
    }
  }

  return sum * multiplier;
};
