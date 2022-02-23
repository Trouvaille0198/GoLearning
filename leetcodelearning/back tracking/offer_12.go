package back_tracking

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var backtracking func(i, j, layer int) bool
	// i,j为board指针 layer为层数 从0开始
	backtracking = func(i, j, layer int) bool {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			// 超出边界或回到了上一层
			return false
		}
		if board[i][j] != word[layer] {
			// 字符不同
			return false
		}
		if layer == len(word)-1 {
			return true
		}
		board[i][j] = ' ' // 用来防止回到上一层
		res := backtracking(i+1, j, layer+1) || backtracking(i-1, j, layer+1) ||
			backtracking(i, j+1, layer+1) || backtracking(i, j-1, layer+1)
		board[i][j] = word[layer]
		return res
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtracking(i, j, 0) {
				return true
			}
		}
	}
	return false
}
