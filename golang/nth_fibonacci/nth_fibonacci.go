package main

import (
	"fmt"
	"reflect"
)

/*

Difficulty: Easy
Category: Recursion / Dynammic Programming
Nth Fibonacci

The Fibonacci sequence is defined as follows: the first number of the sequence is 0, the second number is 1, and the nth number is the sum of the (n - 1)th and (n - 2)th numbers. Write a function that takes in an integer n and returns the nth Fibonacci number.

Important note: the Fibonacci sequence is often defined with its first two numbers as F0 = 0 and F1 = 1. For the purpose of this question, the first Fibonacci number is F0; therefore, getNthFib(1) is equal to F0, getNthFib(2) is equal to F1, etc..

Sample Input #1
n = 2

Sample Output #1
1 // 0, 1

Sample Input #2
n = 6

Sample Output #2
5 // 0, 1, 1, 2, 3, 5

Hint 1

The formula to generate the nth Fibonacci number can be written as follows: F(n) = F(n - 1) + F(n - 2). Think of the case(s) for which this formula doesn't apply (the base case(s)) and try to implement a simple recursive algorithm to find the nth Fibonacci number with this formula.

Hint 2

What are the runtime implications of solving this problem as is described in Hint #1? Can you use memoization (caching) to improve the performance of your algorithm?

Hint 3

Realize that to calculate any single Fibonacci number you only need to have the two previous Fibonacci numbers. Knowing this, can you implement an iterative algorithm to solve this question, storing only the last two Fibonacci numbers at any given time?

Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the input number

Test Cases
{ n: 6, want: 5 },
{ n: 1, want: 0 },
{ n: 2, want: 1 },
{ n: 3, want: 1 },
{ n: 4, want: 2 },
{ n: 5, want: 3 },
{ n: 7, want: 8 },
{ n: 8, want: 13 },
{ n: 9, want: 21 },
{ n: 10, want: 34 },
{ n: 11, want: 55 },
{ n: 12, want: 89 },
{ n: 13, want: 144 },
{ n: 14, want: 233 },
{ n: 15, want: 377 },
{ n: 16, want: 610 },
{ n: 17, want: 987 },
{ n: 18, want: 1597 },
*/

func main() {
  tests := []struct{
    n int
    want int
  }{
    { n: 6, want: 5 },
    { n: 1, want: 0 },
    { n: 2, want: 1 },
    { n: 3, want: 1 },
    { n: 4, want: 2 },
    { n: 5, want: 3 },
    { n: 7, want: 8 },
    { n: 8, want: 13 },
    { n: 9, want: 21 },
    { n: 10, want: 34 },
    { n: 11, want: 55 },
    { n: 12, want: 89 },
    { n: 13, want: 144 },
    { n: 14, want: 233 },
    { n: 15, want: 377 },
    { n: 16, want: 610 },
    { n: 17, want: 987 },
    { n: 18, want: 1597 },
  }

  for _, tc := range tests {
    got := nthFibRecursive(tc.n)
    if !reflect.DeepEqual(tc.want, got) {
      fmt.Printf("FAIL! Expected: %d, Got: %d\n", tc.want, got)
    } else {
      fmt.Printf("PASS! Expected: %d, Got: %d\n", tc.want, got)
    }
  }
}

