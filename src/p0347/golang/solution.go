func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	a := make([]Pair, 0)
	for k, v := range m {
		a = append(a, Pair{value: k, count: v})
	}
	h := make(Heap, 0, k)
	for _, v := range a {
		if len(h) < k {
			(&h).Push(v)
		} else if v.count > h[0].count {
			(&h).Pop()
			(&h).Push(v)
		}
	}
	ans := make([]int, 0, k)
	for len(h) > 0 {
		x := h.Pop()
		ans = append(ans, x.value)
	}

	return ans
}

type Pair struct {
	value int
	count int
}

type Heap []Pair

func (h *Heap) siftUp(i int) {
	for i > 0 && (*h)[(i-1)>>1].count > (*h)[i].count {
		(*h)[(i-1)>>1], (*h)[i] = (*h)[i], (*h)[(i-1)>>1]
		i = (i - 1) >> 1
	}
}

func (h *Heap) siftDown(i int) {
	smallest := i
	left, right := 2*i+1, 2*i+2
	if left < len(*h) && (*h)[left].count < (*h)[smallest].count {
		smallest = left
	}
	if right < len(*h) && (*h)[right].count < (*h)[smallest].count {
		smallest = right
	}
	if i != smallest {
		(*h)[smallest], (*h)[i] = (*h)[i], (*h)[smallest]
		h.siftDown(smallest)
	}
}

func (h *Heap) Push(x Pair) {
	*h = append(*h, x)
	h.siftUp(len(*h) - 1)
}

func (h *Heap) Pop() Pair {
	x := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	*h = (*h)[0 : len(*h)-1]
	h.siftDown(0)
	return x
}