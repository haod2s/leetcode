func robotSim(commands []int, obstacles [][]int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	to := 0
	ans := 0
	stop := make(map[int]map[int]bool)
	for _, v := range obstacles {
		if _, ok := stop[v[0]]; !ok {
			stop[v[0]] = map[int]bool{}
		}
		stop[v[0]][v[1]] = true
	}
	for x, y, i := 0, 0, 0; i < len(commands); i++ {
		if commands[i] >= 1 && commands[i] <= 9 {
			for j := 1; j <= commands[i]; j++ {
				k := to % 4
				if k < 0 {
					k += 4
				}
				m := x + directions[k][0]
				n := y + directions[k][1]
				if stop[m][n] {
					break
				}
				if m*m+n*n > ans {
					ans = m*m + n*n
				}
				x, y = m, n
			}
		} else if commands[i] == -2 {
			to -= 1
		} else if commands[i] == -1 {
			to += 1
		}
	}
	return ans
}