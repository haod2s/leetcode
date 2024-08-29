package main

import "strconv"

func calculate(s string) int {
	nums := make([]int, 0)
	ops := make([]byte, 0)

	calc := func() int {
		x, y := nums[len(nums)-2], nums[len(nums)-1]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		switch op {
		case '+':
			return x + y
		case '-':
			return x - y
		case '*':
			return x * y
		}

		return x / y
	}

	for i := 0; i < len(s); i++ {
		if isdigit(s[i]) {
			j := i + 1
			for ; j < len(s) && isdigit(s[j]); j++ {
			}
			x, _ := strconv.Atoi(s[i:j])
			nums = append(nums, x)
			i = j
		}
		if i >= len(s) {
			break
		}
		if s[i] == '(' {
			ops = append(ops, s[i])
		}
		if s[i] == ')' {
			for ops[len(ops)-1] != '(' {
				nums = append(nums, calc())
			}
			ops = ops[:len(ops)-1]
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			for len(ops) > 0 && priorTo(ops[len(ops)-1], s[i]) {
				nums = append(nums, calc())
			}
			ops = append(ops, s[i])
		}
	}

	for len(ops) > 0 {
		nums = append(nums, calc())
	}

	return nums[0]
}

func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func priorTo(top, cur byte) bool {
	return cur == '+' || cur == '-' || top == '*' || top == '/'
}
