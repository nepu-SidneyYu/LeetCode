package twoPreSum

func countSubmatrices(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	sumMatrix := make([][]int, m)
	sumMatrix[0] = make([]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		sumMatrix[i+1] = make([]int, n)
		for j := 0; j < n; j++ {
			sumMatrix[i+1][j+1] = sumMatrix[i+1][j] + sumMatrix[i][j+1] - sumMatrix[i][j]
			if sumMatrix[i][j] <= k {
				ans++
			}
		}
	}
	return ans
}
