package main

func findKthPositive(arr []int, k int) int {
	n := len(arr)
	left, right := 0, n-1
	remain := k
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if arr[mid]-mid-1 < k {
			remain = k - (arr[mid] - mid - 1)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return mid - remain - 1
}
