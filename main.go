package main

func setZeroes1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	mark := make([][]bool, m)
	for i := 0; i < len(matrix); i++ {
		mark[i] = make([]bool, len(matrix[0]))
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

func main() {
	setZeroes1([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
}
