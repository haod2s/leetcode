func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	q := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				q = append(q, []int{i, j})
			}
		}
	}
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	groups := [][][]int{q}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dirs {
				x, y := p[0]+d[0], p[1]+d[1]
				if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] != 1 && dis[x][y] == 0 {
					q = append(q, []int{x, y})
					dis[x][y] = len(groups)
				}
			}
		}
		groups = append(groups, q)
	}
	dsu := genDSU(n * n)
	for ans := len(groups) - 2; ans > 0; ans-- {
		for _, p := range groups[ans] {
			for _, d := range dirs {
				x, y := p[0]+d[0], p[1]+d[1]
				if x >= 0 && x < n && y >= 0 && y < n && dis[x][y] >= dis[p[0]][p[1]] {
					dsu.merge(p[0]*n+p[1], x*n+y)
				}
			}
		}
		if dsu.find(0) == dsu.find(n*n-1) {
			return ans
		}
	}
	return 0
}

type dsu struct {
	pa []int
}

func genDSU(n int) *dsu {
	d := &dsu{
		pa: make([]int, n),
	}
	for n > 0 {
		d.pa[n-1] = -1
		n--
	}
	return d
}

func (d *dsu) find(x int) int {
	if d.pa[x] < 0 {
		return x
	}
	d.pa[x] = (*d).find(d.pa[x])
	return d.pa[x]
}

func (d *dsu) merge(x, y int) {
	p, q := d.find(x), d.find(y)
	if p == q {
		return
	}
	if d.pa[p] > d.pa[q] {
		d.pa[q] += d.pa[p]
		d.pa[p] = q
	} else {
		d.pa[p] += d.pa[q]
		d.pa[q] = p
	}
}