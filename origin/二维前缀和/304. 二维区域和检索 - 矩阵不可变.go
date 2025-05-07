package 二维前缀和

type NumMatrix struct {
	SumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sumMatrix := make([][]int, m+1)
	sumMatrix[0] = make([]int, n+1)
	for i, _ := range matrix {
		sumMatrix[i+1] = make([]int, n+1) //每个都需要初始化
		for j, _ := range matrix[i] {
			sumMatrix[i+1][j+1] = sumMatrix[i+1][j] + sumMatrix[i][j+1] - sumMatrix[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{
		SumMatrix: sumMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.SumMatrix[row2+1][col2+1] - this.SumMatrix[row2+1][col1] - this.SumMatrix[row1][col2+1] + this.SumMatrix[row1][col1]
}
