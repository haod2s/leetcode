package main

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	cases := []struct {
		cost   []int
		wanted int
	}{
		{
			cost:   []int{10, 15, 20},
			wanted: 15,
		},
	}
	for _, c := range cases {
		if ans := minCostClimbingStairs(c.cost); ans != c.wanted {
			t.Errorf("want %d, but %d", c.wanted, ans)
		}
	}
}

func Test_solutionMemo(t *testing.T) {
	cases := []struct {
		cost   []int
		wanted int
	}{
		{
			cost:   []int{10, 15, 20},
			wanted: 15,
		},
		{
			cost:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			wanted: 6,
		},
		{
			cost:   []int{10, 20},
			wanted: 10,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionMemo(c.cost); ans != c.wanted {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}

func Test_solutionDP(t *testing.T) {
	cases := []struct {
		cost   []int
		wanted int
	}{
		{
			cost:   []int{10, 15, 20},
			wanted: 15,
		},
		{
			cost:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			wanted: 6,
		},
		{
			cost:   []int{10, 20},
			wanted: 10,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionDP(c.cost); ans != c.wanted {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}
