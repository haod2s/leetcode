func permute(nums []int) [][]int {
	N := len(nums)
	output := nums
	outputs := make([][]int, 0)
	return backtrack(N, 0, output, outputs)
}

func backtrack(N, i int, output []int, outputs [][]int) [][]int {
	if N == i {
		var t []int
		t = append(t, output...)
		outputs = append(outputs, t)
	}
	for j := i; j < N; j++ {
		output[i], output[j] = output[j], output[i]
		outputs = backtrack(N, i+1, output, outputs)
		output[i], output[j] = output[j], output[i]
	}
	return outputs
}
