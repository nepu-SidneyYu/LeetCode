package main

import "sort"

func threeSum(nums []int) [][]int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num] = 1
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result := 0 - nums[i] - nums[j]
			if mp[result] != 0 {
				if i > 0 && j > 1 {
					if nums[i-1] != nums[i] && nums[j] != nums[j-1] {
						res = append(res, []int{nums[i], nums[j], result})
						mp[result]--
					}
				}
			}
		}
	}
	return res
}
