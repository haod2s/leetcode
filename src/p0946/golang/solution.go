func validateStackSequences(pushed []int, popped []int) bool {
	st := make([]int, 0)
	for i, j := 0, 0; i < len(pushed); i++ {
		st = append(st, pushed[i])
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}