package main

import (
	"fmt"
	"reflect"
)

func getExpressionSums(num string) int32 {
}

func main() {
	tests := []struct {
		nums string
		want int32
	}{
		{nums: "123", want: 168},
		{nums: "12", want: 15},
		{nums: "111", want: 308},
		{nums: "1", want: 1},
		{nums: "202", want: 606},
		{nums: "10", want: 11},
		{nums: "999", want: 2994},
		{nums: "101", want: 202},
		{nums: "56", want: 67},
		{nums: "2345", want: 27255},
	}

	for _, tc := range tests {
		got := getExpressionSums(tc.nums)

		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAILED! nums: %s, got: %d, want: %d\n", tc.nums, got, tc.want)
		} else {
			fmt.Printf("PASSED! nums: %s, got: %d, want: %d\n", tc.nums, got, tc.want)
		}
	}
}
