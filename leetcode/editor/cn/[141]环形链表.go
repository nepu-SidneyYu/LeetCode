
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
if head == nil || head.Next == nil {
return false
}
slow, fast := head, head.Next

for fast != nil && fast.Next != nil {
if fast == slow {
return true
}
fast = fast.Next.Next
slow = slow.Next
}
return false
}
//leetcode submit region end(Prohibit modification and deletion)
