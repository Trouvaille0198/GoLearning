package back_tracking

// DFS
func maxDepth(root *TreeNode) int {
	var backtracking func(root *TreeNode, layer int) int
	backtracking = func(root *TreeNode, layer int) int {
		// 返回以root为节点的树的深度
		if root == nil {
			return 0
		}
		l1 := backtracking(root.Left, layer+1)
		l2 := backtracking(root.Right, layer+1)
		if l1 > l2 {
			return 1 + l1
		} else {
			return 1 + l2
		}
	}
	return backtracking(root, 0)
}

// BFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			tmp := queue[0]
			queue = queue[1:] // 出队
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		res++
	}
	return res
}
