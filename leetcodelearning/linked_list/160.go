package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		a = a.Next
		b = b.Next

		if a == nil && b == nil {
			// 不相交的两个节点走到尽头
			return nil
		}
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
	}
	return a
}
