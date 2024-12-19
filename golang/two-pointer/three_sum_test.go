package main

import (
	"reflect"
	"testing"
)

func TestFindSumOfThree(t *testing.T) {
	tests := []struct {
		target int
		arr    []int
		want   bool
	}{
		{target: -1, arr: []int{1, -1, 0}, want: false},
		{target: 10, arr: []int{3, 7, 1, 2, 8, 4, 5}, want: true},
		{target: 21, arr: []int{3, 7, 1, 2, 8, 4, 5}, want: false},
		{target: -8, arr: []int{-1, 2, 1, -4, 5, -3}, want: true},
		{target: 0, arr: []int{-1, 2, 1, -4, 5, -3}, want: true},
		{target: 20, arr: []int{3, 7, 1, 2, 8, 4, 5}, want: true},
	}

	for _, tc := range tests {
		got := FindSumOfThree(tc.arr, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Test target: %v, expected: %v, got: %v", tc.target, tc.want, got)
		}
	}
}
