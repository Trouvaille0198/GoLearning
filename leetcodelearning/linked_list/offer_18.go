package linked_list

// 递归
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	head.Next = deleteNode(head.Next, val)
	return head
}

// 双指针
func deleteNode2(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	preNode := head
	curNode := head
	for curNode != nil {
		if curNode.Val == val {
			preNode.Next = curNode.Next
			return head
		}
		preNode = curNode
		curNode = curNode.Next
	}
	return head
}
