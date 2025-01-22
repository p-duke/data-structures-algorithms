package main

import (
	"fmt"
	"reflect"
)

// Implement below

// TESTS
func main() {
	got := sheet.getCell("A1")
	want := "55"
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("FAILED! got: %s, want: %s\n", got, want)
	} else {
		fmt.Printf("PASSED! got: %s, want: %s\n", got, want)
	}

	sheet.setCell("B1", "20")
	got = sheet.getCell("A1")
	want = "65"
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("FAILED! got: %s, want: %s\n", got, want)
	} else {
		fmt.Printf("PASSED! got: %s, want: %s\n", got, want)
	}
}
