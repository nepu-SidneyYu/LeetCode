package tree

func maxDepth1(root *TreeNode) int {
	Maxdepth := 0
	var inorder func(root *TreeNode, depth int)
	inorder = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		if depth > Maxdepth {
			Maxdepth = depth
		}
		inorder(root.Left, depth)
		inorder(root.Right, depth)
	}
	inorder(root, 0)
	return Maxdepth
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
