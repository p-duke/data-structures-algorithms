// O(n) time | O(1) space
function isPalindrome(string) {
	let j = string.length - 1;
	for (let i = 0; i < string.length; i++) {
		if (string[i] !== string[j]) {
			return false;
		}
		j -= 1;
	}
	
	return true;
}

// Do not edit the line below.
exports.isPalindrome = isPalindrome;

