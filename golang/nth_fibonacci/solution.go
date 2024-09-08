
// O(n) time | O(n) space
func nthFibMemoized(n int) int {
	if n <= 1 {
		return n
	}

	// Create a memoization table to store Fibonacci numbers
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

// O(2^n) exponential time - We need to re-calculate fib(n) several times
// O(n) space
func nthFibRecursive(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return nthFibRecursive(n-1) + nthFibRecursive(n-2)
	}
}

// O(n) time | O(1) space
func nthFibIterative(n int) int {
	// Problem
	// Input represents the nth number in the fib sequence we need to find
	// Return that number in the sequence
	// Note that n of 1 is 0 fib

	// Solution
	// Fib = (n - 1) + (n - 2)
	// Iterate n times
	// Keep track of fibSum
	// At end of loop return fibSum
	if n == 1 {
		return 0
	}

	var p1 int = 0
	var p2 int = 1
	var fibSum int = p1 + p2

	for i := 2; i < n; i++ {
		fibSum = p1 + p2
		p1 = p2
		p2 = fibSum
	}

	return fibSum
}

// Iterative
func fib(n int) int {
	sum := 0
	prev := 1
	for i := 0; i < n; i++ {
        temp := sum
        sum = sum + prev
        prev = temp
	}

	return sum
}

