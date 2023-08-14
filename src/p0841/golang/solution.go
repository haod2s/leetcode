func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	vis := make([]bool, n)
	vis[0] = true
	n--
	q := make([]int, 0)
	for _, k := range rooms[0] {
		q = append(q, k)
	}
	for len(q) > 0 {
		k := q[0]
		q = q[1:]
		if !vis[k] {
			n--
			vis[k] = true
			for _, v := range rooms[k] {
				q = append(q, v)
			}
		}
	}
	return n == 0
}