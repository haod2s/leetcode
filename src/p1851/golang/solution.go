func minInterval(intervals [][]int, queries []int) []int {
	ans := make([]int, len(queries))
	indexes := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return queries[indexes[i]] < queries[indexes[j]]
	})
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := make(minHeap, 0)
	heap.Init(&h)
	for i, j := 0, 0; i < len(indexes); i++ {
		ans[indexes[i]] = -1
		for ; j < len(intervals); j++ {
			if queries[indexes[i]] < intervals[j][0] {
				break
			}
			heap.Push(&h, interval{
				left:  intervals[j][0],
				right: intervals[j][1],
			})
		}
		for h.Len() > 0 {
			x := h[0]
			if x.right >= queries[indexes[i]] {
				ans[indexes[i]] = x.right - x.left + 1
				break
			}
			heap.Pop(&h)
		}
	}
	return ans
}

type interval struct {
	left  int
	right int
}

type minHeap []interval

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Less(i, j int) bool { return h[i].right-h[i].left < h[j].right-h[j].left }

func (h *minHeap) Push(v any) { *h = append(*h, v.(interval)) }
func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}