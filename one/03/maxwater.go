package one

import (
	"fmt"
	"strings"
)

func maxVowels(s string, k int) int {
	str := "aeiou"
	c := 0
	for m := 0; m < k; m++ {
		if strings.Contains(str, string(s[m])) {
			c += 1
		}
	}
	count := c
	i, j := 1, k
	for j < len(s) {
		if strings.Contains(str, string(s[i-1])) {
			c -= 1
		}
		if strings.Contains(str, string(s[j])) {
			c += 1
		}
		count = Max(c, count)
		i++
		j++
	}
	return count
}
func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	res := 0
	l := len(flowerbed)
	mark := make([]int, l+2)
	for i, a := range flowerbed {
		mark[i+1] = a
	}
	for i := 1; i < l+2; i++ {
		if mark[i-1] == 0 && mark[i+1] == 0 && mark[i] == 0 {
			res++
			mark[i] = 1
		}
	}
	if res >= n {
		return true
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = Min(leftMin[i-1], nums[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}
