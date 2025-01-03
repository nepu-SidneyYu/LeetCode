package half

func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if citations[mid] < n-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n - left
}
