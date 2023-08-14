func numMagicSquaresInside(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			if magic(grid[i][j], grid[i][j+1], grid[i][j+2],
				grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
				grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2]) {
				ans++
			}
		}
	}
	return ans
}

func magic(nums ...int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if v < 1 || v > 9 {
			return false
		}
		m[v] += 1
		if m[v] > 1 {
			return false
		}
	}
	return (nums[0]+nums[1]+nums[2] == nums[3]+nums[4]+nums[5] &&
		nums[3]+nums[4]+nums[5] == nums[6]+nums[7]+nums[8] &&
		nums[0]+nums[3]+nums[6] == nums[1]+nums[4]+nums[7] &&
		nums[1]+nums[4]+nums[7] == nums[2]+nums[5]+nums[8] &&
		nums[0]+nums[4]+nums[8] == nums[2]+nums[4]+nums[6] &&
		nums[0]+nums[1]+nums[2] == nums[0]+nums[3]+nums[6] &&
		nums[0]+nums[1]+nums[2] == nums[0]+nums[4]+nums[8])
}