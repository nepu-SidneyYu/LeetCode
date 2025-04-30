package half

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	l, r := 0, n-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == nums[mid^1] { //^运算 与1^运算偶数+1，奇数-1，mid为偶数时说明单一元素出现在左边
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
