package InterviewQuestions

import "fmt"

//func permute(nums []int) [][]int {
//	n := len(nums)
//	res := make([][]int, 0)
//	mark := make([]bool, n)
//	path := make([]int, n)
//	var dfs func(i int)
//	dfs = func(i int) {
//		if i == n {
//			res = append(res, append([]int{}, path...))
//		}
//		for j, _ := range mark {
//			if !mark[j] {
//				path[i] = nums[j]
//				mark[j] = true
//				dfs(i + 1)
//				mark[j] = false
//			}
//		}
//	}
//	dfs(0)
//	return res
//}

func main() {
	nums := []int{3, 2, 1}
	res := permute1(nums)
	fmt.Printf("%#v\n", res)
}

func permute1(nums []int) [][]int {
	res := make([][]int, 0)
	mark := make(map[int]bool)
	for _, num := range nums {
		mark[num] = false
	}
	path := make([]int, 0)
	for _, num := range nums {
		path = append(path, num)
		mark[num] = true
		res = append(res, dfs1(path, nums, mark)...)
		mark[num] = false
		path = make([]int, 0)
	}
	return res
}

func dfs1(path []int, nums []int, mark map[int]bool) (res [][]int) {

	if len(path) == len(nums) {
		res = append(res, path)
		return res
	}
	for _, num := range nums {
		if !mark[num] {
			temp := make([]int, len(path))
			copy(temp, path)
			temp = append(temp, num)
			path = append(path, num)
			mark[num] = true
			res = append(res, dfs1(temp, nums, mark)...)
			//path = path[0 : len(path)-1]
			mark[num] = false
		}
	}
	return res
}

func dfs2(path []int, nums []int, mark map[int]bool) (res [][]int) {

	if len(path) == len(nums) {
		res = append(res, path)
		return res
	}
	for _, num := range nums {
		if !mark[num] {
			//temp := make([]int, len(path))
			//copy(temp, path)
			//temp = append(temp, num)
			path = append(path, num)
			mark[num] = true
			res = append(res, dfs2(path, nums, mark)...)
			path = path[0 : len(path)-1]
			mark[num] = false
		}
	}
	return res
}
