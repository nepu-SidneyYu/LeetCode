package leetcodehot100

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := make(map[int]int)
	cur := head
	for cur != nil {
		count[cur.Val]++
		cur = cur.Next
	}

	dummy := &ListNode{Next: head}
	//tmp := dummy
	pre := dummy
	cur = head
	for cur != nil {
		if count[cur.Val] == 1 {
			pre.Next = cur
			pre = pre.Next
		}
		cur = cur.Next
	}
	pre.Next = nil
	return dummy.Next
}
