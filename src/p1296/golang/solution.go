func isPossibleDivide(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	keys := make([]int, 0, len(m))
	for num, _ := range m {
		keys = append(keys, num)
	}
	sort.Ints(keys)
	total := len(nums)
	for total > 0 {
		x := -1
		for i := 0; i < len(keys); i++ {
			if m[keys[i]] > 0 {
				x = keys[i]
				break
			}
		}
		for d := 0; d < k; d++ {
			if m[x+d] <= 0 {
				return false
			}
			m[x+d]--
			total--
		}
	}
	return true
}