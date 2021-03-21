// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
// 进阶：
//
// 一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个仅使用常量空间的解决方案吗？
//
//
// 示例 1：
//
//
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
// 示例 2：
//
//
// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
//
// 提示：
//
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -231 <= matrix[i][j] <= 231 - 1
// 通过次数86,547提交次数149,486

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// mark position
	pos := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				pos = append(pos, []int{i, j})
			}
		}
	}

	for _, p := range pos {
		i, j := p[0], p[1]
		up := i - 1
		for up >= 0 {
			matrix[up][j] = 0
			up--
		}
		down := i + 1
		for down < m {
			matrix[down][j] = 0
			down++
		}
		left := j - 1
		for left >= 0 {
			matrix[i][left] = 0
			left--
		}
		right := j + 1
		for right < n {
			matrix[i][right] = 0
			right++
		}
	}
}