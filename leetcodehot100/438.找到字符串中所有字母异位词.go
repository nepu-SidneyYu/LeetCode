package leetcodehot100

// 双map效率高的，能够通过数组直接比较里面的元素的多少，切片无法直接比较数组是否相等
func findAnagrams1(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	n := len(p)
	i, j := 0, n-1
	var mark, markp [26]int
	for k := i; k < j; k++ {
		mark[s[k]-'a']++
	}
	for _, a := range p {
		markp[a-'a']++
	}
	res := make([]int, 0)
	for j < len(s) {
		mark[s[j]-'a']++
		if mark == markp {
			res = append(res, i)
		}
		mark[s[i]-'a']--
		i++
		j++
	}
	return res
}

// 效率低
func findAnagrams0(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	n := len(p)
	i, j := 0, n-1
	mark := make([]int, 26)
	markp := make([]int, 26)
	//var mark, markp [26]int
	for k := i; k < j; k++ {
		mark[s[k]-'a']++
	}
	for _, a := range p {
		markp[a-'a']++
	}
	res := make([]int, 0)
	for j < len(s) {
		mark[s[j]-'a']++
		isNotAnagrams := false
		for k := i; k <= j; k++ {
			if mark[s[k]-'a'] != markp[s[k]-'a'] {
				isNotAnagrams = true
				break
			}
		}
		if !isNotAnagrams {
			res = append(res, i)
		}
		//if mark == markp {
		//	res = append(res, i)
		//}
		mark[s[i]-'a']--
		i++
		j++
	}
	return res
}
