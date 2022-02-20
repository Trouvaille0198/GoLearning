package tree

// 从头结点开始 是不是子树
func isSubStructureHelp(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil {
		// B不空A空
		return false
	}

	if A.Val == B.Val {
		return isSubStructureHelp(A.Left, B.Left) && isSubStructureHelp(A.Right, B.Right)
	}
	return false
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSubStructureHelp(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
