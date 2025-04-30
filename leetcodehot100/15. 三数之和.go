package leetcodehot100

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	for first := 0; first < n; first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := n - 1
		target := -nums[first]
		for second < third {
			if nums[second]+nums[third] < target {
				second++
			} else if nums[second]+nums[third] > target {
				third--
			} else {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				second++
				third--
			}
		}
	}
	return res
}
