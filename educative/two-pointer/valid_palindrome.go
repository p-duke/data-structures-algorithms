package main

import "fmt"
import "strings"

// Write a function that takes a string s as input and checks whether it’s a palindrome or not.
// A phrase, word or sequence is a palindrome that reads the same backwards as forwards.
// The string won’t have any spaces and will only consist of ASCII characters
// Use the Two Pointer pattern
func main() {
	var tests []string = []string{
		"kayaK",
		"hello",
		"RACEACAR",
		"A",
		"ABCDABCD",
		"DCBAABCD",
		"ABCBA",
	}

	for _, v := range tests {
		var result bool = isPalindrome(v)
		fmt.Println(v, "is palindrome?", result)
	}
}

// O(n) time | O(1) space
// The space complexity of the “isPalindrome” function is O(1) because it uses a constant amount of extra space regardless of the size of the input string
func isPalindrome(inputString string) bool {
	newInputString := strings.ToLower(inputString)
	// iterate over the inputString from outside to middle
	// if the two elements don't match return false
	// else return true

	var pTwo int = len(newInputString) - 1
	for i := 0; i < len(newInputString); i++ {
		charOne := string(newInputString[i])
		charTwo := string(newInputString[pTwo])

		if charOne != charTwo {
			return false
		}
		pTwo -= 1
	}

	return true
}

func isPalindromeTwo(inputString string) bool {
	left := 0
	right := len(inputString) - 1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}
	return true
}
