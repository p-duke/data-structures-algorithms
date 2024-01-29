package main

import (
  "fmt"
  "reflect"
)

// Time complexity O(n) | Space complexity O(1)
// Note: Use a two pointer solution is not viable when there's thousands of prices - too long execution
func maxProfit(prices []int) int {
  profit := 0
  minPrice := prices[0]

  for i := 0; i < len(prices); i++ {
    if prices[i] < minPrice {
      minPrice = prices[i]
    } else if (prices[i] - minPrice) > profit {
      profit = prices[i] - minPrice
    }
  }
  return profit
}



func main() {
  tests := []struct{
    input []int
    want int
  }{
    { input: []int{7,1,5,3,6,4}, want: 5 },
    { input: []int{7,6,4,3,1}, want: 0 },
  }

  for _, tc := range tests {
    got := maxProfit(tc.input)
    if !reflect.DeepEqual(got, tc.want) {
      fmt.Printf("FAILED! got: %d, want: %d\n", got, tc.want)
    } else {
      fmt.Printf("PASSED! got: %d, want: %d\n", got, tc.want)
    }
  }
}
