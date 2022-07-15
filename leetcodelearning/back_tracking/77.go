package back_tracking

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	set := make([]int, 0)
	var dfs func(startIndex, layer int)
	dfs = func(startIndex, layer int) {
		if layer == k {
			res = append(res, append([]int{}, set...))
			return
		}
		if n-startIndex-1 < k-len(set) {
			// 剪枝 如果当前能够提供的元素数量小于还需要的元素数量 那就剪去
			return
		}
		for i := startIndex; i < n; i++ {
			set = append(set, i+1)
			dfs(i+1, layer+1)
			set = set[:len(set)-1]
		}
	}
	dfs(0, 0)
	return res
}
