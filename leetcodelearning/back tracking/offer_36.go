package back_tracking

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var inOrder func(root *TreeNode)
	var pre, head *TreeNode
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)

		if pre == nil {
			// 遍历到头节点
			head = root
		}
		root.Left = pre
		pre.Right = root

		pre = root
		inOrder(root.Right)
	}
	inOrder(root)
	// pre最终指向尾结点
	pre.Right = head
	head.Left = pre
	return head
}
