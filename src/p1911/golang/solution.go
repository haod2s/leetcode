func maxAlternatingSum(nums []int) int64 {
	max := func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}
	even, odd := int64(nums[0]), int64(0)
	for i := 1; i < len(nums); i++ {
		x := even
		even = max(even, odd+int64(nums[i]))
		odd = max(x-int64(nums[i]), odd)
	}
	return even
}