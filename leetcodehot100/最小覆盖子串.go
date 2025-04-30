package leetcodehot100

func minWindow(s string, t string) string {
	cnt, ori := make(map[byte]int), make(map[byte]int)
	for i, _ := range t {
		cnt[t[i]]++
	}

	n := len(s)
	sl, sr := 0, 0

	left, right := -1, n

	check := func(cnt, roi map[byte]int) bool {
		for k, v := range cnt {
			if roi[k] < v {
				return false
			}
		}
		return true
	}

	for sr < n {
		ori[s[sr]]++
		for check(cnt, ori) && sl <= sr {
			if sr-sl < right-left {
				left, right = sl, sr
			}
			ori[s[sl]]--
			sl++
		}
		sr++
	}
	if left == -1 {
		return ""
	}
	return s[left : right+1]
}
