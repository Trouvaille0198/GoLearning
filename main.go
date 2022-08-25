package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getFamily(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res1, res2 int
	if root.Left != nil && root.Right != nil && root.Val > 0 && root.Left.Val > 0 && root.Right.Val > 0 {
		root.Val, root.Left.Val, root.Right.Val = -root.Val, -root.Left.Val, -root.Right.Val
		leftSum := getFamily(root.Left)
		rightSum := getFamily(root.Right)
		res1 = leftSum + rightSum + 1
		// 还原
		root.Val, root.Left.Val, root.Right.Val = -root.Val, -root.Left.Val, -root.Right.Val
	}
	leftSum := getFamily(root.Left)
	rightSum := getFamily(root.Right)
	res2 = leftSum + rightSum

	return max(res1, res2)
}

// func getNodeCnt(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return 1 + getNodeCnt(root.Left) + getNodeCnt(root.Right)
// }

// func getFamily(root *TreeNode) int {
// 	tmpRes, res := 0, 0
// 	layer := 0
// 	nodeCnt := getNodeCnt(root)
// 	var dfs func(root *TreeNode)
// 	dfs = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		layer++
// 		if layer == nodeCnt {
// 			if tmpRes > res {
// 				res = tmpRes
// 			}
// 			tmpRes, layer = 0, 0
// 			return
// 		}

// 		if root.Left != nil && root.Right != nil {
// 			if root.Val > 0 && root.Left.Val > 0 && root.Right.Val > 0 {
// 				root.Val, root.Left.Val, root.Right.Val = -root.Val, -root.Left.Val, -root.Right.Val
// 				tmpRes++
// 				dfs(root.Left)
// 				dfs(root.Right)
// 				// 还原
// 				root.Val, root.Left.Val, root.Right.Val = -root.Val, -root.Left.Val, -root.Right.Val
// 				tmpRes--
// 			}
// 		}
// 		dfs(root.Left)
// 		dfs(root.Right)
// 	}
// 	dfs(root)
// 	return res
// }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
