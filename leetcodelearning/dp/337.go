package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	f, g := map[*TreeNode]int{}, map[*TreeNode]int{}
	var postOrder func(root *TreeNode)
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)

		f[root] = root.Val + g[root.Left] + g[root.Right]
		g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])
	}
	postOrder(root)
	return max(f[root], g[root])
}

type robSum struct {
	f, g int
}

func rob2(root *TreeNode) int {
	var postOrder func(root *TreeNode) robSum
	postOrder = func(root *TreeNode) robSum {
		if root == nil {
			return robSum{0, 0}
		}
		robSumLeft := postOrder(root.Left)
		robSumRight := postOrder(root.Right)

		f := root.Val + robSumLeft.g + robSumRight.g
		g := max(robSumLeft.f, robSumLeft.g) + max(robSumRight.f, robSumRight.g)

		return robSum{f, g}
	}
	robSum := postOrder(root)
	return max(robSum.f, robSum.g)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
