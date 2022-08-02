package others

type Node struct {
	Key, Val int
	Next     *Node
	Prev     *Node
}

type LinkList struct {
	Head, Tail *Node
}

func NewLinkList() *LinkList {
	res := LinkList{Head: &Node{}, Tail: &Node{}}
	res.Head.Next = res.Tail
	res.Tail.Prev = res.Head
	return &res
}

func (l *LinkList) AddToHead(n *Node) {
	l.Head.Next.Prev = n
	n.Next = l.Head.Next
	l.Head.Next = n
	n.Prev = l.Head
}

func (l *LinkList) MoveToHead(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	l.AddToHead(n)

}

func (l *LinkList) RemoveTail() int {
	n := l.Tail.Prev
	l.Tail.Prev = n.Prev
	n.Prev.Next = l.Tail
	return n.Key
}

type LRUCache struct {
	HashMap  map[int]*Node
	LinkList *LinkList
	Capacity int
	Size     int
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{HashMap: make(map[int]*Node), LinkList: NewLinkList(), Capacity: capacity}
	return res
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.HashMap[key]; !ok {
		return -1
	} else {
		node := this.HashMap[key]
		// 命中，移至最前
		this.LinkList.MoveToHead(node)
		return node.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.HashMap[key]; !ok {
		// 不存在
		node := Node{Key: key, Val: value}
		this.HashMap[key] = &node
		this.LinkList.AddToHead(&node)
		this.Size++
		if this.Size > this.Capacity {
			// 溢出，淘汰最末
			removedKey := this.LinkList.RemoveTail()
			delete(this.HashMap, removedKey)
			this.Size--
		}
	} else {
		// 存在，修改其值
		node := this.HashMap[key]
		node.Val = value
		// 命中，移至最前
		this.LinkList.MoveToHead(node)
	}
}
