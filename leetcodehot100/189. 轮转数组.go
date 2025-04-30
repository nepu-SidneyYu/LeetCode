package leetcodehot100

import "fmt"

func rotate1(nums []int, k int) {
	n := len(nums)
	rotateNum := k % n

	temp := make([]int, 0)
	temp = append(temp, nums[n-rotateNum:n]...)
	temp = append(temp, nums[:n-rotateNum]...)
	fmt.Println(temp)
	for i, _ := range temp {
		nums[i] = temp[i]
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	mod := k % n
	reserve(nums, 0, n-1)
	reserve(nums, 0, mod-1)
	reserve(nums, mod, n-1)
}
func reserve(nums []int, left, right int) {
	i, j := left, right
	for i < j {
		nums[j], nums[i] = nums[i], nums[j]
		i++
		j--
	}
}
