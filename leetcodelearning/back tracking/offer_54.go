package back_tracking

func kthLargest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	var i, res int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		i++
		if i == k {
			res = root.Val
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}
