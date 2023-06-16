// Solution 1
function getNthFib(n) {
	// Fib sequence:
	// 1  2  3  4  5  6  7  8
	// 0, 1, 1, 2, 3, 5, 8, 13
	if (n === 1) return 0;

	let i = 2;
	let p1 = 0;
	let p2 = 1;
	let sum = p1 + p2;

	while (i < n) {
		sum = p1 + p2;
		p1 = p2;
		p2 = sum;
		i += 1;
	}
	
	return sum;
}

// Do not edit the line below.
exports.getNthFib = getNthFib;

// Solution 2
// time - 2^n because we need to re-calculate fib(n) several times
function getNthFib(n) {
	if (n === 2) {
		return 1;
	} else if (n === 1) {
		return 0;
	} else {
		return getNthFib(n - 1) + getNthFib(n - 2);
	}
}

// Do not edit the line below.
exports.getNthFib = getNthFib;

// Memoized
// O(n) time | O(n) space
function getNthFib(n,  memoized = { 1: 0, 2: 1 }) {
	if (typeof memoized[n] === 'number') {
		return memoized[n];
	} else {
		memoized[n] = getNthFib(n - 1, memoized) + getNthFib(n - 2, memoized);
		return memoized[n];
	}
}

// Do not edit the line below.
exports.getNthFib = getNthFib;

// Solution 3
// O(n) time | O(1) space
function getNthFib(n) {
	let lastTwo = [0, 1];
	let counter = 3;

	while (counter <= n) {
		let nextFib = lastTwo[0] + lastTwo[1];
		lastTwo[0] = lastTwo[1];
		lastTwo[1] = nextFib;
		counter += 1;
	}
	
	return n > 1 ? lastTwo[1] : lastTwo[0];
}

// Do not edit the line below.
exports.getNthFib = getNthFib;
