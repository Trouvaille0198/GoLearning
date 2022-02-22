package tree

import "container/list"

// 非递归
func postorderTraversal(root *TreeNode) []int {
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

		result = append([]int{curNode.Val}, result...) // 将结果放在最前面

		if curNode.Left != nil {
			stack.PushBack(curNode.Left)
		}
		if curNode.Right != nil {
			stack.PushBack(curNode.Right)
		}
	}
	return result
}

// 非递归
func postorderTraversal2(root *TreeNode) []int {
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

		if curNode.Left != nil {
			stack.PushBack(curNode.Left)
		}
		if curNode.Right != nil {
			stack.PushBack(curNode.Right)
		}
	}
	reverse := func(nums []int) []int {
		i, j := 0, len(nums)-1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		return nums
	}
	return reverse(result)
}

// 递归
func postorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0)
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		help(root.Left)
		help(root.Right)
		result = append(result, root.Val)
	}
	help(root)
	return result
}
