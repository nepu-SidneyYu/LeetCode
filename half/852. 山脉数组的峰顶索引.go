package half

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	l, r := 0, n-1
	mid := 0
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < arr[mid-1] {
			r = mid
		} else if arr[mid] < arr[mid+1] {
			l = mid
		} else {
			return mid
		}
	}
	return mid
}
