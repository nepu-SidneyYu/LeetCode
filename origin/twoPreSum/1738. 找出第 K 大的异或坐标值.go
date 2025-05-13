package twoPreSum

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	orMatrix := make([][]int, m+1)
	orMatrix[0] = make([]int, n+1)
	nums := make([]int, 0)
	for i, _ := range matrix {
		orMatrix[i+1] = make([]int, n+1)
		for j, _ := range matrix[i] {
			orMatrix[i+1][j+1] = orMatrix[i+1][j] ^ orMatrix[i][j+1] ^ orMatrix[i][j] ^ matrix[i][j]
		}
		nums = append(nums, orMatrix[i+1][1:]...)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
