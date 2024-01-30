package main

import (
	"fmt"
	"reflect"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
*/

func isValid(s string) bool {
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, char := range s {
		// Opening
		if _, ok := pairs[char]; ok {
			stack = append(stack, char)
        // Closing invalid
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != char {
            return false
        // Closing valid
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}


func main() {
	tests := []struct{
		input string
		want bool
	}{
		{ input: "()", want: true },
		{ input: "()[]{}", want: true },
		{ input: "(]", want: false },
		{ input: "{(([]))}", want: true },
	}

	for _, tc := range tests {
		got := isValid(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAILED! got: %t, want: %t\n", got, tc.want)
		} else {
			fmt.Printf("PASSED! got: %t, want: %t\n", got, tc.want)
		}
	}

}

