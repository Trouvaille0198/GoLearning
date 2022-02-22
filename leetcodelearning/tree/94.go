package tree

import "container/list"

// 非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	result := make([]int, 0)
	curNode := root // 指针
	for stack.Len() != 0 || curNode != nil {
		if curNode != nil {
			stack.PushBack(curNode)
			curNode = curNode.Left
		} else {
			// 左边没有了
			curNode = stack.Back().Value.(*TreeNode) // 出栈
			stack.Remove(stack.Back())

			result = append(result, curNode.Val)
			curNode = curNode.Right
		}
	}
	return result
}

// 递归
func inorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		help(root.Left)
		result = append(result, root.Val)
		help(root.Right)
	}
	help(root)
	return result
}
