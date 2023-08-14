func longestMountain(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}
	ans := 0
	for k := 1; k < n-1; k++ {
		i, j := k-1, k+1
		if arr[k] <= arr[i] || arr[k] <= arr[j] {
			continue
		}
		for ; i > 0; i-- {
			if arr[i-1] >= arr[i] {
				break
			}
		}
		for ; j < n-1; j++ {
			if arr[j] <= arr[j+1] {
				break
			}
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
	}
	return ans
}