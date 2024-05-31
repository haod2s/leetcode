package main

import "testing"

func Test_solutionMemo(t *testing.T) {
	cases := []struct {
		low, high, zero, one, wanted int
	}{
		{
			low:    3,
			high:   3,
			zero:   1,
			one:    1,
			wanted: 8,
		},
		{
			low:    2,
			high:   3,
			zero:   1,
			one:    2,
			wanted: 5,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := solutionMemo(c.low, c.high, c.zero, c.one); ans != c.wanted {
				t.Errorf("want %d, but %d", c.wanted, ans)
			}
		})
	}
}
