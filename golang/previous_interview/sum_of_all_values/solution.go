package main

import (
	"strconv"
)

// TODO: This passes the first test but not the others. Need to complete.
func getExpressionSums(nums string) int32 {
	incrementalTotal := make([]int32, 0)
	runningTotal := 0
	for i := 0; i < len(nums); i++ {
		leftOp, _ := strconv.Atoi(string(nums[0:i+1]))
		rightOp, _ := strconv.Atoi(string(nums[i+1:]))

		sum := leftOp + rightOp
		if len(incrementalTotal) == 0 {
			incrementalTotal = append(incrementalTotal, int32(sum))
		} else {
			lastSum := incrementalTotal[len(incrementalTotal)-1] // pop last value
			incrementalTotal = incrementalTotal[:len(incrementalTotal)-1] // clear the stack
			incrementalTotal = append(incrementalTotal, int32(sum) + lastSum)
		}

		singleVal, _ := strconv.Atoi(string(nums[i]))
		runningTotal += singleVal
	}

	incrementalSum := incrementalTotal[len(incrementalTotal)-1] // pop last value
	return (incrementalSum + int32(runningTotal)) % int32(1000000007)
}

