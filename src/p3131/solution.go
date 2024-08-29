package main

import "math"

func addedInteger(nums1 []int, nums2 []int) int {
	maxInSlice := func(s []int) int {
		ans := math.MinInt
		for _, v := range s {
			ans = max(ans, v)
		}

		return ans
	}

	return maxInSlice(nums2) - maxInSlice(nums1)
}
