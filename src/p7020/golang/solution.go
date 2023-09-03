func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for x := low; x <= high; x++ {
		nums := split(x)
		if len(nums)%2 == 0 {
			i := len(nums) - 1
			j := i / 2
			front := 0
			for ; i > j; i-- {
				front += nums[i]
			}
			back := 0
			for ; j >= 0; j-- {
				back += nums[j]
			}
			if front == back {
				ans++
			}
		}
	}
	return ans
}

func split(n int) []int {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	return nums
}