package tree

// 队列实现层序遍历
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		result = append(result, curNode.Val)
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}
	return result
}
