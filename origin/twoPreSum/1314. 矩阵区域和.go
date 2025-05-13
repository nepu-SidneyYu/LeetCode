package twoPreSum

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	sumMatrix := make([][]int, m+1)
	sumMatrix[0] = make([]int, n+1)
	BlockSum := make([][]int, m)
	for i, _ := range mat {
		sumMatrix[i+1] = make([]int, n+1)
		BlockSum[i] = make([]int, n)
		for j, _ := range mat[i] {
			sumMatrix[i+1][j+1] = sumMatrix[i+1][j] + sumMatrix[i][j+1] - sumMatrix[i][j] + mat[i][j]
		}
	}
	row1, col1, row2, col2 := 0, 0, 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row1, row2, col1, col2 = max(0, i-k), min(m, i+k+1), max(0, j-k), min(n, j+k+1)
			BlockSum[i][j] = sumMatrix[row2][col2] - sumMatrix[row1][col2] - sumMatrix[row2][col1] + sumMatrix[row1][col1]
		}
	}
	return BlockSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
