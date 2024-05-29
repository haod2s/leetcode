package main

import "testing"

func Test_solutionMemo(t *testing.T) {
	cases := []struct {
		n      int
		wanted int
	}{
		{
			n:      3,
			wanted: 3,
		},
	}
	for _, c := range cases {
		if ans := solutionMemo(c.n); ans != c.wanted {
			t.Errorf("want %d, but %d", c.wanted, ans)
		}
	}
}

func Test_solutionDP(t *testing.T) {
	cases := []struct {
		n      int
		wanted int
	}{
		{
			n:      3,
			wanted: 3,
		},
	}
	for _, c := range cases {
		if ans := solutionDP(c.n); ans != c.wanted {
			t.Errorf("want %d, but %d", c.wanted, ans)
		}
	}
}
