package main

import (
	pre "Leetcode/origin/twoPreSum"
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {
	pre := pre.MatrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)
	fmt.Println(pre)
}
