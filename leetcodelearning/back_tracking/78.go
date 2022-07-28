package back_tracking

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(subset []int, layer int)
	dfs = func(subset []int, layer int) {
		if layer == len(nums) {
			res = append(res, append([]int{}, subset...))
			return
		}
		dfs(subset, layer+1)
		dfs(append(subset, nums[layer]), layer+1)
	}
	dfs([]int{}, 0)
	return res
}

func subsets2(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(subset []int, startIndex int)
	dfs = func(subset []int, startIndex int) {
		res = append(res, append([]int{}, subset...))
		if startIndex == len(nums) {
			// 终止条件 可以不加
			return
		}
		for i := startIndex; i < len(nums); i++ {
			subset = append(subset, nums[i])
			dfs(subset, i+1)
			subset = subset[:len(subset)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
