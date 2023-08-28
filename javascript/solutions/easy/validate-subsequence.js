// Best - O(n) time | O(n) space
function isValidSubsequence(array, sequence) {
  let seqIdx = 0;
  for (let i = 0; i < array.length; i++) {
    if (array[i] === sequence[seqIdx]) {
      seqIdx += 1;
    }
  }
  return seqIdx === sequence.length;
}

// Good - O(n) time | O(n) space
function isValidSubsequence(array, sequence) {
	if (sequence.length > array.length) return false;
	var match = [];
	var samePos = false;
	for (var i = 0; i < array.length; i++) {
		if (sequence.includes(array[i]) && (match.length < sequence.length)) {
			match.push(array[i]);
			samePos = sequence.indexOf(array[i]) === match.indexOf(array[i]);
			if (!samePos) return false;
		} 
	}
	
	return match.length === sequence.length;
}

// Do not edit the line below.
exports.isValidSubsequence = isValidSubsequence;

