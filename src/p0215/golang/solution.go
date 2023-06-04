func findKthLargest(nums []int, k int) int {
	h := make(MinHeap, 0, k)
	for i := 0; i < k; i++ {
		h = append(h, nums[i])
	}
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		if nums[i] >= h[0] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}

	return h[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}