func maxSubarraySumCircular(nums []int) int {
	maxSum, minSum, maxSub, minSub, total, maxItem := nums[0], nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		maxSub += nums[i]
		minSub += nums[i]
		if maxSub < nums[i] {
			maxSub = nums[i]
		}
		if maxSub > maxSum {
			maxSum = maxSub
		}
		if minSub > nums[i] {
			minSub = nums[i]
		}
		if minSub < minSum {
			minSum = minSub
		}
		if maxItem < nums[i] {
			maxItem = nums[i]
		}
	}
	if maxItem < 0 && total == minSum {
		return maxItem
	}
	if maxSum < total-minSum {
		return total - minSum
	}
	return maxSum
}
