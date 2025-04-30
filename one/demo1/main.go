package main

func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)
	for _, con := range connections {
		x, y := con[0], con[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], -x)
	}
	vis := make([]bool, n)
	count := 0
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for _, tar := range graph[from] {
			if tar < 0 && !vis[-tar] {
				dfs(-tar)
			} else if tar > 0 && !vis[tar] {
				count++
				dfs(tar)
			}
		}
	}
	return count
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	//p, q := 0, n-1
	var dfs func([]int, int, int, int)
	dfs = func(nums []int, l, r, k int) {
		if l >= r {
			return
		}
		i, j := l-1, r+1
		index := nums[l+r>>1]
		for i < j {
			for i++; nums[i] > index && i < n; i++ {
			}
			for j--; nums[j] < index && j >= 0; j-- {
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if j-l+1 >= k {
			dfs(nums, l, j, k)
		} else {
			dfs(nums, j+1, r, k-(j-l+1))
		}
	}
	dfs(nums, 0, n-1, k)
	return nums[k-1]
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//nums := make([]int, 0)
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	//if j >= 0 {
	//	nums1[:k+1]=nums2[:j+1]
	//}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

//func removeElement(nums []int, val int) int {
//
//}

//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4

func search(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n-1
	for i <= j {
		mid := i + j>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

func nextGreatestLetter(letters []byte, target byte) byte {
	i, j := 0, len(letters)
	for i <= j {
		mid := (i + j) >> 1
		if letters[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if i >= len(letters) {
		return letters[0]
	} else {
		return letters[i]
	}
}

// 输入：grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
// 输出：8
// 解释：矩阵中共有8个负数。
func countNegatives(grid [][]int) int {
	res := 0
	for _, g := range grid {
		n := len(g)
		i, j := 0, n-1
		for i <= j {
			mid := (i + j) >> 1
			if g[mid] < 0 {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
		res += n - i
	}
	return res
}

func searchRange(nums []int, target int) []int {
	return []int{}
}
