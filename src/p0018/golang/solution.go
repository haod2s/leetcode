func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return ans
	}
	sort.Ints(nums)
	for p := 0; p < n-3; p++ {
		if p > 0 && nums[p] == nums[p-1] {
			continue
		}
		for q := p + 1; q < n-2; q++ {
			if q-1 > p && nums[q] == nums[q-1] {
				continue
			}
			i, j := q+1, n-1
			for i < j {
				x := int64(nums[p]) + int64(nums[q]) + int64(nums[i]) + int64(nums[j])
				if x == int64(target) {
					ans = append(ans, []int{nums[p], nums[q], nums[i], nums[j]})
					for j--; j > i && nums[j+1] == nums[j]; j-- {
					}
					for i++; i < j && nums[i] == nums[i-1]; i++ {
					}
				} else if x > int64(target) {
					j--
				} else {
					i++
				}
			}
		}
	}
	return ans
}