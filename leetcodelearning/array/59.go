package array

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	count := 1
	for layer := 0; layer <= n-2; layer++ {
		// 顺时针填空 左闭右开 中心特例
		if layer == n-layer-1 {
			// 中心
			res[layer][layer] = n * n
			break
		}
		// l2r
		for col := layer; col < n-layer-1; col++ {
			res[layer][col] = count
			count++
		}
		// u2d
		for row := layer; row < n-layer-1; row++ {
			res[row][n-layer-1] = count
			count++
		}
		// r2l
		for col := n - layer - 1; col > layer; col-- {
			res[n-layer-1][col] = count
			count++
		}
		// d2u
		for row := n - layer - 1; row > layer; row-- {
			res[row][layer] = count
			count++
		}
	}
	return res
}
