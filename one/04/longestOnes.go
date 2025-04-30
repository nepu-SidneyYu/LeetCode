package one

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func sum(nums []int) int {
	res := 0
	for _, a := range nums {
		res += a
	}
	return res
}

func pivotIndex(nums []int) int {
	Sum := sum(nums)
	leftsum := 0
	for i := 0; i < len(nums); i++ {

		if leftsum == Sum-nums[i]-leftsum {
			return i
		}
		leftsum += nums[i]
	}
	return -1
}

func main() {
	//var nums []int = {1,1,1,0,0,0,1,1,1,1,0}
	//fmt.Println(longestOnes(nums,2))
}
