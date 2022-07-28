package tree

import "math"

// 递归
func isValidBST(root *TreeNode) bool {
	var help func(root *TreeNode, minVal, maxVal int) bool

	help = func(root *TreeNode, minVal, maxVal int) bool {
		if root == nil {
			return true
		}

		if root.Val <= minVal || root.Val >= maxVal {
			return false
		}

		return help(root.Left, minVal, root.Val) && help(root.Right, root.Val, maxVal)
	}

	return help(root, math.MinInt64, math.MaxInt64)
}

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	lastNodeVal := math.MinInt64
	curNode := root

	for len(stack) != 0 || curNode != nil {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else {
			curNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 跟中序遍历的上一个节点值比较
			if curNode.Val <= lastNodeVal {
				return false
			}
			lastNodeVal = curNode.Val
			curNode = curNode.Right
		}
	}
	return true
}
