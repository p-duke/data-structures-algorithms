package main

import (
	"fmt"
	"reflect"
)

// Write a function that takes a string as input and checks whether it can be a valid palindrome by removing at most one character from it.
// O(n) time | O(1) space
func validPalindromeTwo(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}

		left++
		right--
	}

	return true
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	tests := []struct {
		assert string
		want   bool
	}{
		{assert: "madame", want: true},
		{assert: "dead", want: true},
		{assert: "abca", want: true},
		{assert: "tebbem", want: false},
		{assert: "eeccccbebaeeabebccceea", want: false},
		{assert: "ognfjhgbjhzkqhzadmgqbwqsktzqwjexqvzjsopolnmvnymbbzoofzbbmynvmnloposjzvqxejwqztksqwbqgmdazhqkzhjbghjfno", want: false},
		{assert: "abcdedadedecba", want: true},
		{assert: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", want: true},
	}

	for _, tc := range tests {
		got := validPalindromeTwo(tc.assert)
		if !reflect.DeepEqual(tc.want, got) {
			fmt.Printf("FAIL! Assert: %s, Expected: %t --- Got: %t\n", tc.assert, tc.want, got)
		} else {
			fmt.Printf("PASS Assert: %s, Expected: %t --- Got: %t\n", tc.assert, tc.want, got)
		}
	}
}
