func rotate(matrix [][]int) {
	horizontalFlip(matrix)
	diagonalFlip(matrix)
}

func horizontalFlip(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N/2; i++ {
		matrix[i], matrix[N-i-1] = matrix[N-i-1], matrix[i]
	}
}

func diagonalFlip(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}