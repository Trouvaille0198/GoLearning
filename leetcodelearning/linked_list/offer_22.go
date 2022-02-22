package linked_list

// 双指针
func getKthFromEnd(head *ListNode, k int) *ListNode {
	curNode, preNode := head, head
	for curNode != nil {
		curNode = curNode.Next
		if k == 0 {
			preNode = preNode.Next
		} else {
			k--
		}
	}
	return preNode
}
