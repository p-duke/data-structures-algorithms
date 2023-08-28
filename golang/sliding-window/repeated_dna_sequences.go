package main

import (
	"fmt"
	"reflect"
)

// Given a string, s, that represents a DNA sequence, and a number,
// k—return all the contiguous sequences (substrings) of length k
// that occur more than once in the string. The order of the returned
// subsequences does not matter. If no repeated subsequence is found,
// the function should return an empty array.
func findRepeatedSequences(s string, k int) []string {
	var result []string
	sequences := map[string]int{}
	left := 0
	right := k

	for right < len(s) {
		sub := s[left:right]
		if v, ok := sequences[sub]; ok && v < 2 {
			result = append(result, sub)
			sequences[sub] += 1
		} else if _, ok := sequences[sub]; ok {
			sequences[sub] += 1
		} else {
			sequences[sub] = 1
		}

		left++
		right++
	}

	return result
}

func main() {
	tests := []struct {
		sequence string
		k        int
		want     []string
	}{
		{sequence: "AAAAACCCCCAAAAACCCCCC", k: 8, want: []string{"AAAAACCC", "AAAACCCC", "AAACCCCC"}},
		{sequence: "GGGGGGGGGGGGGGGGGGGGGGGGG", k: 12, want: []string{"GGGGGGGGGGGG"}},
		{sequence: "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT", k: 10, want: []string{"TTTTTCCCCC", "CCCCTTTTTT", "CCCCCTTTTT", "CCCCCCTTTT", "TCCCCCCCTT", "TTCCCCCCCT", "CCCCCCCTTT", "TTTCCCCCCC", "TTTTCCCCCC"}},
		// {sequence: "AAAAAACCCCCCCAAAAAAAACCCCCCCTG", k: 10, want: []string{"AAAAAACCCC", "AAACCCCCCC", "AAAACCCCCC", "AAAAACCCCC"}},
		// {sequence: "ATATATATATATATAT", k: 6, want: []string{"ATATAT", "TATATA"}},
	}

	for _, tc := range tests {
		got := findRepeatedSequences(tc.sequence, tc.k)
		if !reflect.DeepEqual(tc.want, got) {
			fmt.Printf("FAIL! Assert: %s, Expected: %v --- Got: %v\n", tc.sequence, tc.want, got)
		} else {
			fmt.Printf("PASS Assert: %s, Expected: %v --- Got: %v\n", tc.sequence, tc.want, got)
		}
	}
}
