package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelp(root.Left, root.Right)
}

// 判断两棵树是否镜像
func isSymmetricHelp(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A == nil || B == nil {
		// 若一个为空一个不为空
		return false
	}

	if A.Val == B.Val && isSymmetricHelp(A.Left, B.Right) && isSymmetricHelp(A.Right, B.Left) {
		return true
	}
	return false
}
