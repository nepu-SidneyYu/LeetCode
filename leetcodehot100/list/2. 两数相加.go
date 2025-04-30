package list

/*
给你两个 非空 的链表，表示两个非负的整数。
它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var head *ListNode
	tail := head
	carry := 0
	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil {
			l1 = l1.Next
			val1 = l1.Val
		}
		if l2 != nil {
			l2 = l2.Next
			val2 = l2.Val
		}
		sum := val1 + val2 + carry
		mod := 0
		carry, mod = sum/10, sum%10
		if head == nil {
			head = &ListNode{Val: mod}
		} else {
			tail.Next = &ListNode{Val: mod}
			tail = tail.Next
		}
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}
