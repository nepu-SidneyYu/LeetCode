package twoPreSum

func numberOfSubmatrices(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ma := make([][][2]int, m+1)
	ma[0] = make([][2]int, n+1)
	ans := 0
	for i, _ := range grid {
		ma[i+1] = make([][2]int, n+1)
		for j, x := range grid[i] {
			ma[i+1][j+1][0] = ma[i+1][j][0] + ma[i][j+1][0] - ma[i][j][0]
			ma[i+1][j+1][1] = ma[i+1][j][1] + ma[i][j+1][1] - ma[i][j][1]
			if x != '.' {
				ma[i+1][j+1][x-'.']++
			}
			if ma[i+1][j+1][0] > 0 && ma[i+1][j+1][0] == ma[i+1][j+1][1] {
				ans++
			}
		}
	}
	return ans
}
