package list

/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	phead := dummy.Next
	pdummy := dummy

	for phead != nil {
		k = k - 1
		for k > 0 {
			if phead == nil {

				return dummy.Next
			}
			phead = phead.Next
			k--
		}
		cur := pdummy.Next
		pre := pdummy
		var next *ListNode
		for cur != nil {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			if cur == phead {
				break
			}
		}
		phead = next
		node := pdummy.Next
		//pdummy = pdummy.Next
		pdummy.Next = cur
		pdummy = node
		node.Next = next
		//cur = next
	}
	return dummy.Next
}
