package back_tracking

func combinationSum3(k int, n int) [][]int {
	// k: 层数
	var res = make([][]int, 0)
	var tmp = make([]int, 0)
	var dfs func(layer, sum, startIndex int)
	dfs = func(layer, sum, startIndex int) {
		if layer == k {
			if sum == n {
				res = append(res, append([]int{}, tmp...))
			}
			return
		}
		for i := startIndex; i <= 9; i++ {
			if sum+i > n {
				continue
			}
			tmp = append(tmp, i)
			dfs(layer+1, sum+i, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, 0, 1)
	return res
}
