package linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur *ListNode = nil, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 跟双指针思想类似的递归
func reverseList(head *ListNode) *ListNode {
	var reverse func(pre, cur *ListNode) *ListNode
	reverse = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}
		after := cur.Next
		cur.Next = pre
		return reverse(cur, after)
	}
	return reverse(nil, head)
}

// 递归2
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head // 翻转头节点与第二个节点的指向
	head.Next = nil
	return last
}
