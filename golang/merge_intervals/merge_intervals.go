package main

import (
	"fmt"
	"reflect"
  "strconv"
)

type Interval struct {
	start  int
	end    int
	closed bool
}

func (i *Interval) intervalInit(start int, end int) {
	i.start = start
	i.end = end

	// By default the interval is closed
	i.closed = true
}

func (i *Interval) newIntervals(itrvl [][]int) {
  for i, v := range 
}

func (i *Interval) setClosed(closed bool) {
	i.closed = closed
}

func (i *Interval) str() string {
	out := ""
	if i.closed {
		out = "[" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + "]"
	} else {
		out = "(" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + ")"
	}
	return out
}


// We are given an array of closed intervals, intervals, where each interval has a start time and an end time. The input array is sorted with respect to the start times of each interval. For example, intervals =
// [ [1,4], [3,6], [7,9] ] is sorted in terms of start times 1, 3 , and 7.
// Your task is to merge the overlapping intervals and return a new output array consisting of only the non-overlapping intervals.
func mergeIntervals(intervals []Interval) []Interval {
	// Replace this placeholder return statement with your code
	return make([]Interval, 0)
}

func main() {
	tests := []struct {
		intervals [][]int
		want     [][]int
	}{
		{intervals: [][]int{{1, 5}, {3, 7}, {4, 6}}, want: [][]int{{1, 7}}},
	}

	for _, tc := range tests {
    intervals := newIntervals(tc.intervals) 

		got := mergeIntervals(tc.intervals)
		if !reflect.DeepEqual(tc.want, got) {
			fmt.Printf("FAIL! Assert: %s, Expected: %v --- Got: %v\n", tc.sequence, tc.want, got)
		} else {
			fmt.Printf("PASS Assert: %s, Expected: %v --- Got: %v\n", tc.sequence, tc.want, got)
		}
	}
}
