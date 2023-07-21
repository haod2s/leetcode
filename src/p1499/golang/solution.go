package golang

import (
	"container/heap"
	"math"
)

func findMaxValueOfEquation(points [][]int, k int) int {
	ans := math.MinInt
	h := maxHeap{{points[0][1] - points[0][0], points[0][0]}}
	heap.Init(&h)
	for i := 1; i < len(points); i++ {
		for h.Len() > 0 && points[i][0]-h[0][1] > k {
			heap.Pop(&h)
		}
		if h.Len() > 0 && points[i][0]+points[i][1]+h[0][0] > ans {
			ans = points[i][0] + points[i][1] + h[0][0]
		}
		heap.Push(&h, []int{points[i][1] - points[i][0], points[i][0]})
	}
	return ans
}

type maxHeap [][]int

func (h maxHeap) Len() int      { return len(h) }
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

func (h *maxHeap) Push(x any) { *h = append(*h, x.([]int)) }
func (h *maxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
