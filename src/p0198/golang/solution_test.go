package main

import "testing"

func Test_solutionMemo(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted int
	}{
		{
			nums:   []int{1},
			wanted: 1,
		},
		{
			nums:   []int{1, 2, 3, 1},
			wanted: 4,
		},
		{
			nums:   []int{2, 7, 9, 3, 1},
			wanted: 12,
		},
		{
			nums:   []int{2, 1},
			wanted: 2,
		},
		{
			nums:   []int{2, 1, 1, 2},
			wanted: 4,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionMemo(c.nums); ans != c.wanted {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}

func Test_solutionDP(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted int
	}{
		{
			nums:   []int{1},
			wanted: 1,
		},
		{
			nums:   []int{1, 2, 3, 1},
			wanted: 4,
		},
		{
			nums:   []int{2, 7, 9, 3, 1},
			wanted: 12,
		},
		{
			nums:   []int{2, 1},
			wanted: 2,
		},
		{
			nums:   []int{2, 1, 1, 2},
			wanted: 4,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionDP(c.nums); ans != c.wanted {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}
