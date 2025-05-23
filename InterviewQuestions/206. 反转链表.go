package InterviewQuestions

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverseList(head.Next)
	head.Next.Next = head //例如4->5->Next = 4   此时就是4->5且5->4成环了
	head.Next = nil       //形成了环形链表，需要破环
	return temp
}

// 迭代
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		head = head.Next
		//head = head.Next
		cur.Next = pre
		pre = cur
		cur = head
	}
	return pre
}
