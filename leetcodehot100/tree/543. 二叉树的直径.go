package tree

func diameterOfBinaryTree(root *TreeNode) int {
	maxLen := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		llen := dfs(node.Left) + 1
		rlen := dfs(node.Right) + 1
		maxLen = max(maxLen, llen+rlen)
		return max(llen, rlen)
	}
	dfs(root)
	return maxLen
}
