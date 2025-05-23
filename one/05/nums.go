package one

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	mp1 := make(map[int]int, 0)
	mp2 := make(map[int]int, 0)
	for _, a := range nums1 {
		mp1[a]++
	}
	for _, a := range nums2 {
		mp2[a]++
	}
	res := make([][]int, 2)
	//res[0]:=make([]int,0)
	for a, _ := range mp1 {
		if _, ok := mp2[a]; !ok {
			res[0] = append(res[0], a)
			//delete(mp2, a)
		}
	}
	for a, _ := range mp2 {
		if _, ok := mp1[a]; !ok {
			res[1] = append(res[1], a)
			//delete(mp1, a)
		}
	}
	return res
}

func uniqueOccurrences(arr []int) bool {
	mp := make(map[int]int)
	for _, num := range arr {
		mp[num]++
	}
	values := make(map[int]int)
	for _, v := range mp {
		if values[v] == 0 {
			values[v]++
		} else {
			return false
		}
	}
	return true
}

func closeStrings(word1 string, word2 string) bool {
	n1 := len(word1)
	n2 := len(word2)
	if n1 != n2 {
		return false
	}
	mp := make(map[byte]int)
	mp1 := make(map[byte]int)
	for _, a := range word1 {
		mp[byte(a)]++
	}
	for _, a := range word2 {
		mp1[byte(a)]++
	}
	for i := 0; i < n1; i++ {
		if mp[byte(word2[i])] == 0 && mp1[byte(word1[i])] == 0 {
			return false
		}
	}
	return true
}

func main() {
	var arr = []int{-130, 21, -154, 159, -44, -126, 165, 68, -126, -126, -126, 128, -94, 165, -30, -44, -39, -94, 21, -130, 68, 68, 128, -130, -39, 181, 68, 68, 68, 139, 139, -39, 21, 21, -39, 68, 128, 131, -126, -154, -30, 165, 21, 159, 181, -39, -126, 131, -94, -44, 131, 128, 21, -44, 128, -94, 183, -94, 131, 139, -44, 128, 21, 181, -44, 131, 128, 131, 21, 68, 181, -44, -126, -130, 131, -190, 131, 181, 165, -94, 165, 165, -30, -154, 68, -39, -44, 165, -39, -126, 68, 68, -130, 68, -94, 181, -44, 131, 21, 183, -44, 21, -39, -130, -39, 131, 21, 165, 165, -126, 165, -44, -94, 68, 68, -94, -126, -126, -30, 181, 165, 68, -44, -39, -94, -126, -126, -30, 68, 181, -44, -94, -126, -44, -94, -30, 131, 165, -190, -130, -94, -94, 181, 128, 181, 181, 181, 139, -130, -94, -130, -130, 139, -130, -90, -154, 181, 165, -30, -154, 165, -190, 159, 165, 139, -126, -44, 131, -44, -190, -126, -130, -94, 128, -154, 68, -130, -130, 68, 21, -44, -30, -126, -126, 131, 159, -190, -126, 181, 139}
	b := uniqueOccurrences(arr)
	fmt.Println(b)
}
