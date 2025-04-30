package half

import "sort"

// func findKthPositive(arr []int, k int) int {
// 	n := len(arr)
// 	left, right := 0, n-1
// 	remain := k
// 	mid := 0
// 	for left <= right {
// 		mid = (left + right) / 2
// 		if arr[mid]-mid-1 < k {
// 			remain = k - (arr[mid] - mid - 1)
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return mid - remain - 1
// }

func findKthPositive1(arr []int, k int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i]-i-1 >= k }) + k
}

func findKthPositive2(arr []int, k int) int {
	n := len(arr)
	left, right := 0, n-1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left + k
}

func findKthPositive(arr []int, k int) int {
	n := len(arr)
	left, right := 0, n-1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left + k
}
