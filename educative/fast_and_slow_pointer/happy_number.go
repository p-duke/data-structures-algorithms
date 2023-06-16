package main

import (
  "fmt"
  "reflect"
  "strings"
  "strconv"
) 

// Educative Solution 
// The time complexity for the solution shown above is O(logn).
// However, since there were two pointers, the cost is O(logn+logn), that is, O(logn) .

// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

// isHappy is the challenge function
func isHappy(num int) bool {
	slowPointer := num
	fastPointer := sumOfSquaredDigits(num)

	for fastPointer != 1 && slowPointer != fastPointer {
		slowPointer = sumOfSquaredDigits(slowPointer)
		fastPointer = sumOfSquaredDigits(sumOfSquaredDigits(fastPointer))
	}

	if fastPointer == 1 {
		return true
	}
	return false
}

// Helper function that calculates the sum of squared digits
func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0{
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}
// ----------------------------------------------------------------------------

// Write an algorithm to determine if a number num is happy.
// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), 
// or it loops endlessly in a cycle which does not include 1
// Those numbers for which this process ends in 1 are happy.
// Return TRUE if num is a happy number, and FALSE if not

// O(n) time | O(1) space
// The built-in Go functions strconv.Itoa and strconv.Atoi have a time complexity of O(n)
// The strings.Split function has a time complexity of O(n)
// The space complexity of the isHappy function is O(1) because it only uses a constant amount of space to store two integers, slow and fast
func isHappy(num int) bool {
  slow := num
  fast := sumOfSquares(num)
  // fmt.Println("slow", slow, "fast", fast)
  for slow < fast {
    if slow == fast {
      return false
    } else {
      slow = sumOfSquares(slow)
      fast = sumOfSquares(sumOfSquares(fast))
    }
    // fmt.Println("slow", slow, "fast", fast)
  }

  if fast == 1 {
    return true
  }

	return false
}

// The sumOfSquaredDigits function has a time complexity of O(log n), where n is the input number.
// This is because the function iterates over each digit in the input number, and the number of digits
// in the input number is proportional to log n.
// The space complexity of the sumOfSquares function is O(log n) because the space used to store the
// digis slice is proportional to the number of digits in the input number, which is log(n) where n is the value of the input number.
func sumOfSquares(num int) int {
  var digi string = strconv.Itoa(num)
  var digis []string = strings.Split(digi, "")
  var result int

  for _, v := range digis {
    convert, err := strconv.Atoi(v)
    if err != nil {
      fmt.Println("String conversion error", err)
    }
    square := convert * convert
    result = result + square
  }

  return result
}


func main() {
	tests := []struct {
		assert int
		want   bool
	}{
    {assert: 1, want: true},
    {assert: 4, want: false},
    {assert: 2147483646, want: false},
    {assert: 19, want: true},
    {assert: 7, want: true},
	}

	for _, tc := range tests {
		got := isHappy(tc.assert)
		if !reflect.DeepEqual(tc.want, got) {
      fmt.Printf("FAIL! Assert: %d, Expected: %t --- Got: %t\n", tc.assert, tc.want, got)
    } else {
      fmt.Printf("PASS Assert: %d, Expected: %t --- Got: %t\n", tc.assert, tc.want, got)
    }
	}
}
