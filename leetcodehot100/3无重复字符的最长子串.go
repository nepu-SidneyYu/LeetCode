package leetcodehot100

func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	i, j := 0, 0
	n := len(b)
	mp := make(map[byte]bool)
	for j < n {

	}
}

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}
