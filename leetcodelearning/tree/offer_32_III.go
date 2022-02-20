package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	curCount := 1
	curLevel := 0
	nextCount := 0
	isAscending := true
	for len(queue) != 0 {
		curNodes := make([]int, 0)
		for curCount > 0 {
			curNode := queue[0]
			queue = queue[1:]
			curNodes = append(curNodes, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
				nextCount++
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
				nextCount++
			}
			curCount--
		}
		if !isAscending {
			// 倒序
			reversedNodes := make([]int, 0)
			for i := len(curNodes) - 1; i >= 0; i-- {
				reversedNodes = append(reversedNodes, curNodes[i])
			}
			result = append(result, reversedNodes)
		} else {
			result = append(result, curNodes)

		}
		curCount = nextCount
		nextCount = 0
		curLevel++
		isAscending = !isAscending
	}
	return result
}
