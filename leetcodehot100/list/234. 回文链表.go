package list

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。
如果是，返回 true ；否则，返回 false 。

输入：head = [1,2,2,1]
输出：true
*/

func isPalindrome1(head *ListNode) bool {
	slow, fast := head, head
	//找到中间节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre *ListNode
	cur := slow
	//翻转后半部分
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	//对比
	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

// 翻转+栈
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	//找到中间节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	stack := make([]*ListNode, 0)
	//放入栈中
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}
	for len(stack) > 0 && head != nil {
		if stack[len(stack)-1].Val != head.Val {
			return false
		}
		stack = stack[0 : len(stack)-1]
		head = head.Next
	}
	return true
}
