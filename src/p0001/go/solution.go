func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	ret := make([]int, 2)
	for i, v := range nums {
		if j, ok := m[v]; ok {
			ret[0], ret[1] = i, j
			break
		}
		m[target-v] = i
	}

	return ret
}