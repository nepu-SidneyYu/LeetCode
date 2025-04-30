package list

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		pre = pre.Next
	}
	pre.Next = slow.Next
	return dummy.Next
}
