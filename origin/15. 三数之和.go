package origin

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[j]+nums[j+1] > 0 {
			break
		}
		if x+nums[k]+nums[k-1] < 0 {
			continue
		}
		for j < k {
			if x+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{x, nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if x+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
