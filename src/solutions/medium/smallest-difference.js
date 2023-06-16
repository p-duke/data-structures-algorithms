// Solution 1
function smallestDifference(arrayOne, arrayTwo) {
	const arrOne = arrayOne.sort((a,b) => a - b);
	const arrTwo = arrayTwo.sort((a,b) => a - b);
	let arrOneIdx = 0;
	let arrTwoIdx = 0;
	let pairs = [];
	let lowest;
	
	while (arrOneIdx < arrOne.length || arrTwoIdx < arrTwo.length) {
		let one = arrOne[arrOneIdx];
		let two = arrTwo[arrTwoIdx];
		let diff =  Math.abs(one-(two));
		
		if (!lowest || diff < lowest) {
			lowest = diff;
			pairs = [one, two];
		}
		
		if ((arrOneIdx < arrOne.length) && one < two) {
			arrOneIdx += 1;
		} else if ((arrTwoIdx < arrTwo.length) && two < one) {
			arrTwoIdx += 1;
		} else {
			return pairs;
		}
	}
	
	return pairs;
}

// Do not edit the line below.
exports.smallestDifference = smallestDifference;

