package leetcodehot100

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLen := 0
	left, right := 0, 0
	mark := [128]bool{}
	for right < n {
		for mark[s[left]] {
			mark[s[left]] = false
			left++
		}
		mark[s[right]] = true
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		right++
	}
	return maxLen
}
