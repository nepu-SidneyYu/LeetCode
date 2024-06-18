package leetcodehot100

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	res := make([]int, n)
	left := 1
	for i := 0; i < n; i++ {
		res[i] = left * right[i]
		left *= nums[i]
	}
	return res
}
