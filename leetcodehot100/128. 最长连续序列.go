package leetcodehot100

import (
	"sort"
)

// 方法一
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func longestConsecutive1(nums []int) int {
	//排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//放入哈希表
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	longestsub := 0
	//index := 0
	sub := 0
	//遍历有序的数组，当没有连续时，
	//通过i直接跳过连续数组的长度(有重复的情况时跳过数组的长度<=此时不连续数字的下标)
	//跳过的目的主要是减少时间复杂度
	for i := 0; i < len(nums); {
		n := nums[i]
		for {
			if mp[n] && i < len(nums) {
				n++
				i++
				sub++
			} else {
				//sub = 0
				break
			}
		}
		longestsub = max(longestsub, sub)
		sub = 0
	}
	return longestsub
}

// 方法二
func longestConsecutive(nums []int) int {
	longestsub := 0
	//放入哈希表
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	//遍历哈希表
	for num := range mp {
		//找到不存在前一个的数
		if !mp[num-1] {
			sub := 1
			curnum := num
			for mp[curnum+1] {
				curnum++
				sub++
			}
			if sub > longestsub {
				longestsub = sub
			}
		}
	}
	return longestsub
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	longestConsecutive(nums)
}
