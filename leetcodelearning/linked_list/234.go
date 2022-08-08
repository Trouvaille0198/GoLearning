package linked_list

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	startOfSecondHalf := reverseList(slow.Next) // 后半部分的头节点 顺便反转一下后半部分

	p, q := head, startOfSecondHalf
	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p, q = p.Next, q.Next
	}
	return true

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur *ListNode = nil, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
	}
	return pre
}
