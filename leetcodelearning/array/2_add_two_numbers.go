package array

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked_list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0} // 虚拟头节点 空空
	var val1, val2 int
	carry := 0      // 进位
	current := head // 浮标

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		} else {
			// 此位空 用0值代替
			val1 = 0
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		} else {
			val2 = 0
		}

		current.Next = &ListNode{Val: (val1 + val2 + carry) % 10}
		current = current.Next

		carry = (val1 + val2 + carry) / 10
	}
	return head.Next // 返回真·头节点
}
