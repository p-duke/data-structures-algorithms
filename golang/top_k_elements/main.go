package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// Step 1: Count the frequency of each element
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Step 2: Group numbers by frequency
	// We use a bucket where the index represents the frequency
	// The value at each index is a list of numbers with that frequency
	maxFreq := len(nums)
	buckets := make([][]int, maxFreq+1)
	for num, freq := range frequencyMap {
		buckets[freq] = append(buckets[freq], num)
	}
	fmt.Println("buckets", buckets)

	// Step 3: Gather the top k frequent elements from the bucket
	result := []int{}
	for i := maxFreq; i >= 0 && len(result) < k; i-- {
		fmt.Printf("buckets i: %d %v\n", i, buckets[i])
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}

	fmt.Println("end result", result)

	// Return only the top k elements
	return result[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // Output: [1, 2]
}
