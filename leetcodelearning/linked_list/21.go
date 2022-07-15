package linked_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	fakeHead := &ListNode{}
	p := fakeHead
	for list1 != nil || list2 != nil {
		if list1 == nil {
			p.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			p.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val > list2.Val {
				p.Next = list2
				list2 = list2.Next
			} else {
				p.Next = list1
				list1 = list1.Next
			}
		}
		p = p.Next
	}
	return fakeHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}
