package list

/*
将两个升序链表合并为一个新的 升序 链表并返回。
新链表是通过拼接给定的两个链表的所有节点组成的。

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var mergedList *ListNode
	if list1.Val < list2.Val {
		mergedList = list1
		list1 = list1.Next
	} else {
		mergedList = list2
		list2 = list2.Next
	}
	head := mergedList //注意保留合并后的头结点
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			mergedList.Next = list1
			list1 = list1.Next
		} else {
			mergedList.Next = list2
			list2 = list2.Next
		}
		mergedList = mergedList.Next
	}
	if list1 != nil {
		mergedList.Next = list1
	}
	if list2 != nil {
		mergedList.Next = list2
	}
	return head
}
