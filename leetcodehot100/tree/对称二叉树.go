package tree

// 递归
func CheckSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && CheckSymmetric(left.Left, right.Right) && CheckSymmetric(left.Right, right.Left)
}
func isSymmetric1(root *TreeNode) bool {
	return CheckSymmetric(root.Left, root.Right)
}

// 迭代
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, root)
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Right, right.Left)
		//queue = append(queue, )
		queue = append(queue, left.Left, right.Right)
	}
	return true
}
