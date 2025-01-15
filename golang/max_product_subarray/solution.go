func maxProduct(nums []int) int {
    globalMax := nums[0]
    currMax := nums[0]
    currMin := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            currMax, currMin = currMin, currMax
        }

        currMax = max(nums[i], currMax * nums[i])
        currMin = min(nums[i], currMin * nums[i])

        if currMax > globalMax {
            globalMax = currMax
        }
    }

    return globalMax
}
