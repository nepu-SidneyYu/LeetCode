package leetcodehot100

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxArea(height []int) int {
	i, j := 0, len(height)
	area := 0
	max := 0
	for i < j {
		area = min(height[i], height[j]) * (j - 1)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		if area >= max {
			max = area
		}
	}
	return max
}
