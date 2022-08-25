package linked_list

func swapPairs(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	fakeHead := &ListNode{Next: head}
	pre := fakeHead
	for pre.Next != nil && pre.Next.Next != nil {
		p, q := pre.Next, pre.Next.Next
		// p q 交换
		pre.Next = q
		p.Next = q.Next
		q.Next = p

		pre = p
	}
	return fakeHead.Next
}

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head.
	}

	p, q := head, head.Next
	p.Next = q.Next
	q.Next = p
	p.Next = swapPairs(p.Next)

	return q
}
