package main

import "testing"

func Test_solutionMemo(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted int
	}{
		{
			nums:   []int{3, 4, 2},
			wanted: 6,
		},
		{
			nums:   []int{2, 2, 3, 3, 3, 4},
			wanted: 9,
		},
		{
			nums:   []int{3, 4, 2, 10, 2, 8, 3, 3},
			wanted: 27,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionMemo(c.nums); c.wanted != ans {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}
