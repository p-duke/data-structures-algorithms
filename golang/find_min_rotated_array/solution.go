func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	min := nums[low]

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] < min {
			min = nums[mid]
		}

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return min
}
