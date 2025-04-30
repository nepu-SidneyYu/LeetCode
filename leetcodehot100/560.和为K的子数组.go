package leetcodehot100

func subarraySum(nums []int, k int) int {
	ans, sum := 0, 0
	count := make(map[int]int)
	count[0] = 1
	for _, num := range nums {
		sum += num
		ans += count[sum-k]
		count[sum]++
	}
	return ans
}
