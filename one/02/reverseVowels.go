package main

import (
	"fmt"
	"strconv"
)

//func reverseVowels(s string) string {
//	t := []byte(s)
//	//mp := make(map[byte]bool)
//	i, j := 0, len(t)-1
//	for i < j {
//		for i < j && !strings.Contains("aeiouAEIOU", string(t[i])) {
//			i++
//		}
//		for i < j && !strings.Contains("aeiouAEIOU", string(t[j])) {
//			j--
//		}
//		if i < j {
//			t[i], t[j] = t[j], t[i]
//		}
//	}
//	return s
//}

func compress(chars []byte) int {
	i, j := 0, 1
	n := len(chars)
	if n <= 1 {
		return n
	}
	count := 1
	var str string
	by := make([]byte, 0)
	var a byte
	for j < n {
		a = chars[i]
		if chars[j] == chars[i] {
			count++
		} else {
			by = append(by, a)
			if count > 1 {
				str = strconv.Itoa(count)
				t := []byte(str)
				by = append(by, t...)
				count = 1

			}
			i = j
		}
		j++
	}
	by = append(by, a)
	if count > 1 {
		str = strconv.Itoa(count)
		t := []byte(str)
		by = append(by, t...)
	}
	for k := 0; k < len(by); k++ {
		chars[k] = by[k]
	}
	return len(by)
}

func main() {
	chars := []byte{'a', 'b', 'c'}
	n := compress(chars)
	for i := 0; i < n; i++ {
		fmt.Println(chars[i])
	}
}
