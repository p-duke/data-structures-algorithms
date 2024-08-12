package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Pointers to the last elements of nums1 and nums2
	p1 := m - 1
	p2 := n - 1
	// Pointer to the last position of the merged array
	p := m + n - 1

	// Merge in reverse order
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are remaining elements in nums2, copy them over
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
