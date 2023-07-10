func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans, minDis := 0, math.MaxInt
	for k := 0; k < len(nums); k++ {
		for i, j := k+1, len(nums)-1; i < j; {
			x := nums[i] + nums[j] + nums[k]
			if x == target {
				return x
			} else if x > target {
				if x-target < minDis {
					minDis = x - target
					ans = x
				}
				j--
			} else {
				if target-x < minDis {
					minDis = target - x
					ans = x
				}
				i++
			}
		}
	}
	return ans
}