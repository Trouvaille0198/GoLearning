package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 分别计算两链表的长度
	i := headA
	aNum := 0
	for i != nil {
		aNum++
		i = i.Next
	}
	i = headB
	bNum := 0
	for i != nil {
		bNum++
		i = i.Next
	}
	var diff int                     // 两链表的差值
	var fastNode, slowNode *ListNode // 要往前跑diff步的链表头 停在原地的链表头
	// 选出长链表的头节点为fastNode
	if aNum > bNum {
		fastNode, slowNode = headA, headB
		diff = aNum - bNum
	} else {
		fastNode, slowNode = headB, headA
		diff = bNum - aNum
	}
	// 往前跑diff步
	for diff != 0 {
		fastNode = fastNode.Next
		diff--
	}
	// 一起跑 直到相遇
	for fastNode != nil && fastNode != slowNode {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}
	return fastNode
}

// 双指针
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

// 哈希
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}
