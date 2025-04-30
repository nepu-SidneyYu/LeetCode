package list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	tmphead := head
	for tmphead != nil {
		mp[tmphead] = &Node{Val: tmphead.Val}
		tmphead = tmphead.Next
	}
	tmphead = head
	for tmphead != nil {
		mp[tmphead].Next = mp[tmphead.Next]
		mp[tmphead].Random = mp[tmphead.Random]
		tmphead = tmphead.Next
	}
	return mp[head]
}
