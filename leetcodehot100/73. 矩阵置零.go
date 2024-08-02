package leetcodehot100

//二维数组标记O(mn)
func setZeroes1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	mark := make([][]bool, m)
	mark[0] = make([]bool, n)
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mark[i][j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mark[i][j] {
				for k := 0; k < m; k++ {
					matrix[k][j] = 0
				}
				for k := 0; k < n; k++ {
					matrix[i][k] = 0
				}
			}
		}
	}
}

//O(m+n)
func setZeroes2(matrix [][]int) {
}

//O(1)
func setZeroes3(matrix [][]int) {

}
