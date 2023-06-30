func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans := nums[0]
	preMin, preMax, curMin, curMax := nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMin = min(nums[i], min(preMin*nums[i], preMax*nums[i]))
		curMax = max(nums[i], max(preMin*nums[i], preMax*nums[i]))
		ans = max(ans, curMax)
		preMin, preMax = curMin, curMax
	}

	return ans
}