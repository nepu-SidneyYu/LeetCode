package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Right), maxDepth(root.Left)) + 1
}

func dfs(root *TreeNode, r []int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		r = append(r, root.Val)
		return
	}
	dfs(root.Left, r)
	dfs(root.Right, r)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1 := []int{}
	r2 := []int{}
	dfs(root1, r1)
	dfs(root2, r2)
	if len(r1) != len(r2) {
		return false
	}
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func dfsright(root *TreeNode, dir bool) int {
//	if root == nil {
//		return 0
//	}
//	return dfsleft(root.Left) + 1
//}
//
//func dfsleft(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return dfsright(root.Right) + 1
//}
//
//func longestZigZag(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	res := longestZigZag(root)
//	return Max(Max(dfsleft(root.Left), dfsright(root.Right)), res)
//}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		left := searchBST(root.Left, val)
		if left != nil {
			return left
		}
	} else {
		right := searchBST(root.Right, val)
		if right != nil {
			return right
		}
	}
	return nil
}

//func deleDfs
//
//func deleteNode(root *TreeNode, key int) *TreeNode {
//
//}

func quickSort(nums []int, l, r int) {

}

func findKthLargest(nums []int, k int) int {
	return 0
}
