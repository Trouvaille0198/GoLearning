package tree

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		targetNode := &TreeNode{Val: val, Left: root}
		return targetNode
	}
	queue := []*TreeNode{root}
	curDepth := 1
	curLevelCnt := 1
	nextLevelCnt := 0
	for curDepth < depth-1 {
		// 出队
		curNode := queue[0]
		queue = queue[1:]
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
			nextLevelCnt++
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
			nextLevelCnt++
		}
		curLevelCnt--
		if curLevelCnt == 0 {
			// 下一层
			curDepth++
			curLevelCnt = nextLevelCnt
			nextLevelCnt = 0
		}
	}

	for _, node := range queue {
		targetNode1, targetNode2 := &TreeNode{Val: val}, &TreeNode{Val: val}
		targetNode1.Left = node.Left
		node.Left = targetNode1

		targetNode2.Right = node.Right
		node.Right = targetNode2
	}
	return root
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	if depth == 2 {
		root.Left = &TreeNode{Val: val, Left: root.Left}
		root.Right = &TreeNode{Val: val, Right: root.Right}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}
