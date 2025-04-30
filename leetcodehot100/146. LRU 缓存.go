package leetcodehot100

type LRUCache struct {
	head, tail *LinkNode
	cache      map[int]*LinkNode
	size       int
	capacity   int
}

type LinkNode struct {
	Key   int
	Value int
	Next  *LinkNode
	Pre   *LinkNode
}

func initLinkNode(key, value int) *LinkNode {
	return &LinkNode{
		Key:   key,
		Value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*LinkNode),
		head:     initLinkNode(0, 0),
		tail:     initLinkNode(0, 0),
	}
	l.head.Next = l.tail
	l.tail.Pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.removeNode(node)
	this.insertToHead(node)
	return node.Value
}

func (this *LRUCache) insertToHead(l *LinkNode) {
	l.Next = this.head.Next
	this.head.Next.Pre = l
	l.Pre = this.head
	this.head.Next = l
}

func (this *LRUCache) removeNode(node *LinkNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		linkNode := initLinkNode(key, value)
		this.insertToHead(linkNode)
		this.size++
		if this.size > this.capacity {
			//删除尾结点
			this.removeNode(this.tail.Pre)
			delete(this.cache, this.tail.Pre.Key)
			this.size--
		}
		this.cache[key] = linkNode
	} else {
		this.removeNode(node)
		this.insertToHead(node)
		node.Value = value
		this.cache[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
