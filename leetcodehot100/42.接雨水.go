package main

import "fmt"

//[0,1,0,2,1,0,1,3,2,1,2,1]

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	res := 0
	n := len(height)
	for i := 1; i < n-1; i++ {
		if height[i-1] > height[i] && height[i] < height[i+1] {
			min := 0
			if height[i-1] < height[i+1] {
				res += (height[i-1] - height[i])
				min = height[i-1]
			} else {
				res += (height[i+1] - height[i])
				min = height[i+1]
			}
			left, right := i-1, i+1
			for left >= 1 && right < n-1 {
				if height[left-1] > height[left] && height[right] < height[right+1] {
					if height[left-1] < height[right+1] {
						res += (height[left-1] - min) * (right + 1 - (left - 1) - 1)
						min = height[left-1]
					} else {
						res += (height[right+1] - min) * (right + 1 - (left - 1) - 1)
						min = height[right-1]
					}
					left--
					right++
				} else {
					break
				}

			}
		}
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}
