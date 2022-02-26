package back_tracking

import "container/list"

// 获取各位数之和
func getSum(num int) (res int) {
	for num != 0 {
		res += num % 10
		num /= 10
	}
	return
}

// DFS
func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m) // 存储已遍历过的坐标
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}
	res := 0
	var backtracking func(i, j int)
	backtracking = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			// 越界
			return
		}
		if visited[i][j] {
			// 已被遍历过
			return
		}
		visited[i][j] = true
		if getSum(i)+getSum(j) <= k {
			res++
			backtracking(i+1, j) // 下
			backtracking(i, j+1) // 右
		}
	}
	backtracking(0, 0)
	return res
}

// BFS
func movingCount2(m int, n int, k int) int {
	visited := make([][]bool, m) // 存储已遍历过的坐标
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	res := 0
	stack := list.New()
	stack.PushBack([2]int{0, 0})

	var curPos [2]int
	var i, j int
	for stack.Len() != 0 {
		// 出栈
		curPos = stack.Front().Value.([2]int)
		stack.Remove(stack.Front())
		i, j = curPos[0], curPos[1]
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			// 越界
			continue
		}
		if visited[i][j] {
			// 已经遍历过
			continue
		}
		visited[i][j] = true
		if getSum(i)+getSum(j) <= k {
			res++
			// 右边和下边入队
			stack.PushBack([2]int{i + 1, j})
			stack.PushBack([2]int{i, j + 1})
		}
	}
	return res
}

// dp
func movingCount3(m int, n int, k int) int {
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	res := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if getSum(i)+getSum(j) <= k && (visited[[2]int{i - 1, j}] || visited[[2]int{i, j - 1}]) {
				res++
				visited[[2]int{i, j}] = true
			}
		}
	}
	return res
}
