package linked_list

// 非常朴素的哈希表
func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = true
		head = head.Next
	}
	return false
}

// 非常高级的双指针
func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
