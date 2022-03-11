package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	fakeHead := &ListNode{Next: head}
	pre, p, q := fakeHead, head, head
	for i := 0; i < n; i++ {
		q = q.Next
	}
	for q != nil {
		pre = p
		p = p.Next
		q = q.Next
	}
	pre.Next = p.Next
	return fakeHead.Next
}
