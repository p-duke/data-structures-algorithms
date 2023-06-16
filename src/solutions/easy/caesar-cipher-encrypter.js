// Solution 1
// O(n) time | O(n) space
function caesarCipherEncryptor(string, key) {
	const alpha = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];
	let final = '';
	for (let i = 0; i < string.length; i++) {
		const letter = string[i];
		const index = alpha.indexOf(letter);
		const cipher = index + key;
		if (cipher > alpha.length - 1) {
			const idx = cipher % alpha.length;
			final += alpha[idx];
		} else {
			final += alpha[cipher];
		}
		
	}
	
	return final;
}

// Do not edit the line below.
exports.caesarCipherEncryptor = caesarCipherEncryptor;

// Solution 2
// O(n) time | O(n) space
function caesarCipherEncryptor(string, key) {
	const newLetters = [];
	const newKey = key % 26;
	for (const letter of string) {
		newLetters.push(getNewLetter(letter, newKey));
	}
	return newLetters.join('');
}

function getNewLetter(letter, key) {
	const newLetterCode = letter.charCodeAt() + key;
	if (newLetterCode <= 122) {
		return String.fromCharCode(newLetterCode);
	} else {
		return String.fromCharCode(96 + (newLetterCode % 122))
	}
}

// Do not edit the line below.
exports.caesarCipherEncryptor = caesarCipherEncryptor;

// Solution 3 (Same as 2)
// O(n) time | O(n) space
function caesarCipherEncryptor(string, key) {
  // Write your code here.
	const newLetters = [];
	const newKey = key % 26;
	const alphabet = 'abcdefghijklmnopqrstuvwxyz'.split('');
	for (const letter of string) {
		newLetters.push(getNewLetter(letter, newKey, alphabet));
	}
	return newLetters.join('');
}

function getNewLetter(letter, key, alphabet) {
	const newLetterCode = alphabet.indexOf(letter) + key;
	return alphabet[newLetterCode % 26];
}

// Do not edit the line below.
exports.caesarCipherEncryptor = caesarCipherEncryptor;

