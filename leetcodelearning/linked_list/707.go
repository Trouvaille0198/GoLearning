package linked_list

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node // 虚拟头节点
	Len  int
}

// 获取下标为index的结点指针
func (this *MyLinkedList) GetNode(index int) *Node {
	if index == -1 {
		return this.Head
	}
	if index > this.Len-1 || index < -1 {
		return nil
	}
	p := this.Head.Next
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p
}

func Constructor() MyLinkedList {
	return MyLinkedList{Head: &Node{}, Len: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.Len-1 {
		return -1
	}
	p := this.GetNode(index)
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	p := &Node{Val: val, Next: this.Head.Next}
	this.Head.Next = p
	this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	p := this.GetNode(this.Len - 1)
	p.Next = &Node{Val: val}
	this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	} else if index > this.Len {
		return
	}
	p := this.GetNode(index - 1)
	q := this.GetNode(index)
	p.Next = &Node{Val: val, Next: q}
	this.Len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Len-1 {
		return
	}
	p := this.GetNode(index - 1)
	q := this.GetNode(index)
	p.Next = q.Next
	this.Len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
