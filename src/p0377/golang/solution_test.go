package main

import "testing"

func Test_solutionMemo(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		wanted int
	}{
		{
			nums:   []int{1, 2, 3},
			target: 4,
			wanted: 7,
		},
		{
			nums:   []int{9},
			target: 3,
			wanted: 0,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionMemo(c.nums, c.target); ans != c.wanted {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}
