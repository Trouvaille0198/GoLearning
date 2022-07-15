package array

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	l, r, t, b := 0, n-1, 0, m-1 // 左右上下
	count := 0
	state := 'r'
	for count < m*n {
		if state == 'r' {
			// l2r
			for col := l; col <= r; col++ {
				res[count] = matrix[t][col]
				count++
			}
			t++
			state = 'd'
		} else if state == 'd' {
			// u2d
			for row := t; row <= b; row++ {
				res[count] = matrix[row][r]
				count++
			}
			r--
			state = 'l'
		} else if state == 'l' {
			// r2l
			for col := r; col >= l; col-- {
				res[count] = matrix[b][col]
				count++
			}
			b--
			state = 'u'
		} else if state == 'u' {
			// d2u
			for row := b; row >= t; row-- {
				res[count] = matrix[row][l]
				count++
			}
			l++
			state = 'r'
		}
	}
	return res
}
