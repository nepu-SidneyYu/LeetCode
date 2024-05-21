package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	mp := make(map[string][]string)
	for _, str := range strs {
		bstr := []byte(str)
		sort.Slice(bstr, func(i, j int) bool {
			return str[i] < str[j]
		})
		sstr := string(bstr)
		mp[sstr] = append(mp[sstr], str)
	}
	for _, s := range mp {
		res = append(res, s)
	}
	return res
}
