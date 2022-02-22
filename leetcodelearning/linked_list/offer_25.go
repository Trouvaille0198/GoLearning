package linked_list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{} // 伪头节点
	curNode := fakeHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			curNode.Next = l2
			break
		}
		if l2 == nil {
			curNode.Next = l1
			break
		}

		if l1.Val > l2.Val {
			curNode.Next = l2
			l2 = l2.Next
		} else {
			curNode.Next = l1
			l1 = l1.Next
		}
		curNode = curNode.Next
	}
	return fakeHead.Next
}
