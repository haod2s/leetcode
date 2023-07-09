func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans = make([][]int, 0)
	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			continue
		}
		for k > 0 && k < len(nums) && nums[k-1] == nums[k] {
			k++
		}
		for i, j := k+1, len(nums)-1; i < j; {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for i+1 < j && nums[i] == nums[i+1] {
					i++
				}
				for i+1 < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j]+nums[k] > 0 {
				j--
			} else {
				i++
			}
		}
	}
	return ans
}