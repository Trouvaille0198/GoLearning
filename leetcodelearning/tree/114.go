package tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var preNode *TreeNode
	for len(stack) != 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if preNode != nil {
			preNode.Left = nil
			preNode.Right = curNode
		}

		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		preNode = curNode
	}
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			// 找到左子树的最右节点pre
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			// 将右子树接到pre右
			pre.Right = root.Right
			// 将左子树变成右子树
			root.Right = root.Left
			root.Left = nil
			// 考虑下一个节点
			root = root.Right
		}
	}
}
