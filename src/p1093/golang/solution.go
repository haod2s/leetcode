func sampleStats(count []int) []float64 {
	stat := make([][]int, 0)
	for i, v := range count {
		if v > 0 {
			stat = append(stat, []int{i, v})
		}
	}
	ans := make([]float64, 5)
	ans[0] = float64(stat[0][0])
	ans[1] = float64(stat[len(stat)-1][0])
	sum, total := 0, 0
	modeNum := 0
	for _, v := range stat {
		sum += v[0] * v[1]
		total += v[1]
		if v[1] > modeNum {
			modeNum = v[1]
			ans[4] = float64(v[0])
		}
	}
	ans[2] = float64(sum) / float64(total)
	pos := 0
	for i, v := range stat {
		if total%2 == 0 && pos+v[1] == total/2 {
			ans[3] = (float64(stat[i][0]) + float64(stat[i+1][0])) / 2
			break
		} else if pos+v[1] > total/2 {
			ans[3] = float64(v[0])
			break
		}
		pos += v[1]
	}
	return ans
}