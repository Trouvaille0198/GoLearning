package array

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	var dfs func(sum, startIndex int)
	dfs = func(sum, startIndex int) {
		if sum == target {
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			res = append(res, combCopy)
			return
		}
		if sum > target {
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			dfs(sum+candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, 0)
	return res
}
