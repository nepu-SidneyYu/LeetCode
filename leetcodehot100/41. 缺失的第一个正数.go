package leetcodehot100

//标记法
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	//先把小于0的数全部置为n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	//把出现的数字的num-1的位置上的数置为负数，相当于i+1出现了
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	//找到第一个不为负数的下表i+1即为没有出现的
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}

	}
	return n + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//方法2：置换法:
func firstMissingPositive(nums []int) int {
	n := len(nums)
	//将在1-N内的数字置换到i-1的位置上,然后就能满足每个位置上正确出现的的数字都是i+1
	for i := 0; i < n; i++ {
		for nums[i] <= n && nums[i] > 0 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	//统计不为i+1的数就行
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
