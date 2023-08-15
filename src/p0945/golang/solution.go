func minIncrementForUnique(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				break
			}
			nums[j]++
			ans++
		}
	}
	return ans
}