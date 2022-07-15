package tree

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
	var bp func(root *TreeNode) (int, bool)
	bp = func(root *TreeNode) (int, bool) {
		// 返回高度 是否是平衡二叉树
		if root == nil {
			return 0, true
		}
		d1, b1 := bp(root.Left)
		d2, b2 := bp(root.Right)

		if b1 && b2 && abs(d1-d2) <= 1 {
			return max(d1, d2) + 1, true
		}
		return max(d1, d2) + 1, false
	}
	_, res := bp(root)
	return res

}
