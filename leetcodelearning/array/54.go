package array

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	layer := 0   // 当前层
	i, j := 0, 0 // 行列指针
	state := 'r'
	count := 0
	for count < m*n {
		if state == 'r' {
			res[count] = matrix[i][j]
			j++
			if j+1 == n-layer {
				state = 'd'
			}
			if j+1 > n-layer {
				// 溢出后的处理
				j = n - layer - 1
				i++
			}
		} else if state == 'd' {
			res[count] = matrix[i][j]
			i++
			if i+1 == m-layer {
				state = 'l'
			}
			if i+1 > m-layer {
				i = m - layer - 1
				j--
			}
		} else if state == 'l' {
			res[count] = matrix[i][j]
			j--
			if j == layer {
				state = 'u'
			}
			if j < layer {
				j = layer
				i--
			}
		} else if state == 'u' {
			res[count] = matrix[i][j]
			i--
			if i == layer {
				state = 'r'
				layer++
				i++
				j++
			}
		}
		count++
		fmt.Println(res)
	}
	return res
}
