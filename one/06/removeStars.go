package main

import (
	"fmt"
	"strconv"
)

func removeStars(s string) string {
	str := make([]byte, 0)
	for _, a := range s {
		if a == '*' {
			str = str[:len(str)-1]
		} else {
			str = append(str, byte(a))
		}
	}
	return string(str)
}

//func asteroidCollision(asteroids []int) []int {
//	if len(asteroids) <= 1 {
//		return asteroids
//	}
//	nums := make([]int, 0)
//	for i := 1; i < len(asteroids); i++ {
//
//	}
//}

func decodeString(s string) string {
	str := []string{}
	di := []int{}
	by := ([]byte)(s)
	count := ""
	ss := ""
	//rst := ""
	t := ""
	for i := 0; i < len(by); i++ {
		if by[i] >= '0' && by[i] <= '9' {
			count += string(by[i])
			continue
		} else if by[i] >= 'a' && by[i] <= 'z' {
			ss += string(by[i])
		} else if by[i] == '[' {
			if ss != "" {
				str = append(str, ss)
			}
			c, _ := strconv.Atoi(count)
			di = append(di, c)
			count = ""
			ss = ""
		} else if by[i] == ']' {
			if ss != "" {
				str = append(str, ss)
			}
			for j := 0; j < di[len(di)-1]; j++ {
				t += str[len(str)-1]
			}
			di = di[:len(di)-1]
			str = str[:len(str)-1]
			if len(di) > 0 {
				str[len(str)-1] += t
			}
			ss = ""
			t = ""
		}
	}
	return t
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

func predictPartyVictory(senate string) string {
	var ran, di []int
	for i, a := range senate {
		if byte(a) == 'R' {
			ran = append(ran, i)
		} else {
			di = append(di, i)
		}
	}
	n := len(senate)
	for len(ran) > 0 && len(di) > 0 {
		if ran[0] < di[0] {
			ran = append(ran, ran[0]+n)
		} else {
			di = append(di, di[0]+n)
		}
		di = di[1:]
		ran = ran[1:]
	}
	if len(ran) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	s := "3[a2[c]]"
	str := decodeString(s)
	fmt.Println(str)
}
