package others

func rotate(matrix [][]int) {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		end := n - 1 - layer
		for i := layer; i < end; i++ {
			a := matrix[layer][i]     // 左上
			b := matrix[i][end]       // 右上
			c := matrix[end][n-1-i]   // 右下
			d := matrix[n-1-i][layer] // 左下

			matrix[layer][i] = d
			matrix[i][end] = a
			matrix[end][n-1-i] = b
			matrix[n-1-i][layer] = c
		}
	}
}
