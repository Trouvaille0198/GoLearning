package linked_list

// 请实现 copyRandomList 函数，复制一个复杂链表。
// 在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 哈希表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	nodeMap := make(map[*Node]*Node)
	newHead := &Node{} // 空头结点

	for p, k := head, newHead; p != nil; p, k = p.Next, k.Next {
		k.Next = &Node{Val: p.Val}
		nodeMap[p] = k.Next
	}
	for p, k := head, newHead; p != nil; p, k = p.Next, k.Next {
		if p.Random != nil {
			k.Next.Random = nodeMap[p.Random]
		}
	}
	return newHead.Next
}

// 原地转换
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return head
	}
	// 复制结点插入到原节点后面
	for p := head; p != nil; p = p.Next.Next {
		p.Next = &Node{Val: p.Val, Next: p.Next}
	}
	// 设置Random指针
	for p := head; p != nil; p = p.Next.Next {
		k := p.Next
		if p.Random != nil {
			k.Random = p.Random.Next
		}
	}
	newHead := head.Next
	// 拆分原节点和新节点
	for p := head; p != nil; p = p.Next {
		k := p.Next
		p.Next = k.Next
		if k.Next != nil {
			k.Next = k.Next.Next
		}
	}
	return newHead
}
