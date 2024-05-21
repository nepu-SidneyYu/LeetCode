package leetcodehot100

import "sort"

type Pair struct {
	Index int
	Val   int
}

func twoSum(nums []int, target int) []int {
	len := len(nums)
	numspair := make([]Pair, len)
	for i, v := range nums {
		numspair[i].Val = v
		numspair[i].Index = i
	}
	sort.Slice(numspair, func(i, j int) bool {
		return numspair[i].Val < numspair[j].Val
	})
	left, right := 0, len-1
	for left < right {
		if numspair[left].Val+numspair[right].Val == target {
			return []int{numspair[left].Index, numspair[right].Index}
		} else if numspair[left].Val+numspair[right].Val < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
