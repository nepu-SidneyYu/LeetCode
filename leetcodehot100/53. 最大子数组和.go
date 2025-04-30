package leetcodehot100

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	Maxsum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		fmt.Println(nums[i])
		if Maxsum < nums[i] {
			Maxsum = nums[i]
		}
	}
	return Maxsum
}
