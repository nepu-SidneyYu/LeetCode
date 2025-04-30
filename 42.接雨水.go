package leetcodehot100

// [0,1,0,2,1,0,1,3,2,1,2,1]
func trap1(height []int) int {
	stack := make([]int, 0)
	sum := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			cursum := (trapmin(height[left], h) - height[mid]) * (i - left - 1)
			sum += cursum
		}
		stack = append(stack, i)
	}
	return sum
}

func trapmin1(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	n := len(height)
	left, right := 0, n-1
	maxL, maxR := height[0], height[n-1]
	ans := 0
	for left < right {
		if height[left] < height[right] {
			maxL = trapmax(maxL, height[left])
			ans += maxL - height[left]
			left++
		} else {
			maxR = trapmax(maxR, height[right])
			ans += trapmax(maxR, height[right])
			right--
		}
	}
	return ans
}

func trapmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
