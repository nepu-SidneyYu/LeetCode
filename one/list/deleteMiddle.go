package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	dummy := &ListNode{Val: 0, Next: head}
	var fast, slow *ListNode = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		dummy = dummy.Next
	}
	dummy.Next = slow.Next
	return head
}

//func oddEvenList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	j, o := head, head.Next
//	ohead := o
//	jhead := j
//	head = head.Next
//	for head != nil && head.Next != nil {
//		j.Next = head.Next
//		j = j.Next
//		o.Next = j.Next
//		o = o.Next
//		head = head.Next.Next
//	}
//	j.Next = ohead
//	return jhead
//}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyj := &ListNode{Val: 0, Next: nil}
	dummyo := &ListNode{Val: 0, Next: nil}
	jhead := dummyj
	ohead := dummyo
	count := 1
	for head != nil {
		if count%2 == 0 {
			dummyo.Next = head
			dummyo = dummyo.Next
		} else {
			dummyj.Next = head
			dummyj = dummyj.Next
		}
		head = head.Next
		count++
	}
	dummyo.Next = nil
	//dummyj.Next=nil
	dummyj.Next = ohead.Next
	return jhead.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dummy *ListNode
	cur, next := head, head.Next
	for cur != nil {
		next = next.Next
		cur.Next = dummy
		dummy = cur
		cur = next
	}
	return dummy
}
