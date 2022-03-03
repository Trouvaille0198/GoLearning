package tree

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	minDepth := math.MaxInt
	var res *TreeNode
	var bp func(root *TreeNode) (pv, qv bool, d int)
	bp = func(root *TreeNode) (pv, qv bool, d int) {
		// 返回以root为头节点的树是否含有p和q 和它的深度
		if root == nil {
			return false, false, 0
		}

		pv1, qv1, d1 := bp(root.Left)
		pv2, qv2, d2 := bp(root.Right)
		d = max(d1, d2) + 1
		if pv1 || pv2 || root.Val == p.Val {
			pv = true
		}
		if qv1 || qv2 || root.Val == q.Val {
			qv = true
		}
		if pv && qv && d < minDepth {
			res = root
			minDepth = d
		}
		return
	}
	bp(root)
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var a = 1
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
