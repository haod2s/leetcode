package main

import "testing"

func Test_solutionDP(t *testing.T) {
	cases := []struct {
		pressedKeys string
		wanted      int
	}{
		{
			pressedKeys: "22233",
			wanted:      8,
		},
		{
			pressedKeys: "222222222222222222222222222222222222",
			wanted:      82876089,
		},
	}
	for _, c := range cases {
		if ans := solutionDP(c.pressedKeys); c.wanted != ans {
			t.Errorf("want %d, but %d", c.wanted, ans)
		}
	}
}
