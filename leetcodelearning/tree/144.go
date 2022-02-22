package tree

import (
	"container/list"
)

// 非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	var curNode *TreeNode
	for stack.Len() != 0 {
		curNode = stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		result = append(result, curNode.Val)
		if curNode.Right != nil {
			stack.PushBack(curNode.Right)
		}
		if curNode.Left != nil {
			stack.PushBack(curNode.Left)
		}
	}
	return result
}

// 非递归
func preorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		help(root.Left)
		help(root.Right)
	}
	help(root)
	return result
}
