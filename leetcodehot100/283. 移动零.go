package leetcodehot100

func moveZeroes(nums []int) {
	i, j := 0, 0
	n := len(nums)
	for j < n {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
