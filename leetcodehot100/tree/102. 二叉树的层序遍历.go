package tree

// 递归
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return res
}

// 队列遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	cur := []*TreeNode{}

	cur = append(cur, root)
	for len(cur) > 0 {
		nxt := []*TreeNode{}
		vals := make([]int, len(cur))
		for i, node := range cur {
			vals[i] = node.Val
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		res = append(res, vals)
	}
	return res
}
