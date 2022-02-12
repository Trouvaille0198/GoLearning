package linked_list

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

// Definition for singly-linked_list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 辅助栈
func reversePrint(head *ListNode) []int {
	var stack []int
	for p := head; p != nil; p = p.Next {
		stack = append(stack, p.Val)
	}
	var result []int
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
	}
	return result
}

// 递归
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint2(head.Next), head.Val)
}
