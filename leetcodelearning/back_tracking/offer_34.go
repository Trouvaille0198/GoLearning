package back_tracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	road := make([]int, 0)
	var backtracking func(root *TreeNode, leftVal int)
	backtracking = func(root *TreeNode, leftVal int) {
		if root == nil {
			return
		}

		road = append(road, root.Val)
		defer func() {
			road = road[:len(road)-1] // 收回最后一步
		}()

		leftVal -= root.Val
		if leftVal == 0 && root.Left == nil && root.Right == nil {
			// 到叶子节点 且路径和等于目标值
			res = append(res, append([]int{}, road...)) // 拷贝 因为road唯一 会不断变化
			return
		}
		backtracking(root.Left, leftVal)
		backtracking(root.Right, leftVal)
	}
	backtracking(root, target)
	return res
}
