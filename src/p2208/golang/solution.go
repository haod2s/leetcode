func halveArray(nums []int) int {
	h := make(maxHeap, 0, len(nums))
	sum := float64(0)
	for _, v := range nums {
		h = append(h, float64(v))
		sum += float64(v)
	}
	heap.Init(&h)
	ans := 0
	for gap := float64(0); gap < sum/2; {
		h[0] /= 2
		gap += h[0]
		ans++
		heap.Fix(&h, 0)
	}
	return ans
}

type maxHeap []float64

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *maxHeap) Push(x any) { *h = append(*h, x.(float64)) }
func (h *maxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}