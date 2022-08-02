package linked_list

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	length := 0
	for n := head; n != nil; n = n.Next {
		length++
	}

	fakeHead := &ListNode{Next: head}
	for subLen := 1; subLen < length; subLen *= 2 {
		prev, cur := fakeHead, fakeHead.Next
		for cur != nil {
			nodeA := cur
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next
			}

			nodeB := cur.Next
			cur.Next = nil // 斩断nodeA
			cur = nodeB
			for i := 1; i < subLen && cur != nil; i++ {
				cur = cur.Next
			}

			var nextNode *ListNode
			if cur != nil {
				nextNode = cur.Next
				cur.Next = nil // 斩断nodeB
			}

			prev.Next = merge(nodeA, nodeB)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = nextNode
		}
	}
	return fakeHead.Next
}

func merge(a, b *ListNode) *ListNode {
	fakeHead := &ListNode{}
	p := fakeHead
	for {
		if a == nil && b == nil {
			break
		} else if a == nil {
			p.Next = b
			b = b.Next
		} else if b == nil {
			p.Next = a
			a = a.Next
		} else {
			if a.Val > b.Val {
				p.Next = b
				b = b.Next
			} else {
				p.Next = a
				a = a.Next
			}
		}
		p = p.Next
	}
	return fakeHead.Next
}
