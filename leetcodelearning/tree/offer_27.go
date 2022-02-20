package tree

// 递归
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)

	root.Left = right
	root.Right = left
	return root
}
