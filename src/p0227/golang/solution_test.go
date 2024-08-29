package main

import "testing"

func Test_calculate(t *testing.T) {
	cases := []struct {
		formula string
		wanted  int
	}{
		{
			formula: "3+2*3",
			wanted:  9,
		},
		{
			formula: " 3/2 ",
			wanted:  1,
		},
		{
			formula: " 3+5 / 2",
			wanted:  5,
		},
		{
			formula: "3*2+4",
			wanted:  10,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := calculate(c.formula); ans != c.wanted {
				t.Errorf("wanted %d, but %d", c.wanted, ans)
			}
		})
	}
}
