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

/**
 * @param {string} s
 * @return {boolean}
 */
// O(n) time | O(1) space
var isPalindrome = function(s) {
   const newS = s.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();

   if (newS.length === 0) {
       return true;
   }

   let right = newS.length - 1;
   for (let left = 0; left < newS.length; left++) {
       if (!(newS[left] === newS[right])) {
           console.log("returned false", newS[left], newS[right]);
           return false;
       }
       right--;
   }

   return true;
};
