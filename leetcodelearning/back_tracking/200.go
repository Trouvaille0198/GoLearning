package back_tracking

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}
	var explore func(x, y int)
	explore = func(x, y int) {
		if x < 0 || x > rows-1 || y < 0 || y > cols-1 || grid[x][y] == '0' || visited[x][y] {
			return
		}
		visited[x][y] = true
		explore(x+1, y)
		explore(x, y+1)
		explore(x-1, y)
		explore(x, y-1)
	}
	count := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == '0' || visited[x][y] == true {
				continue
			}
			count++
			explore(x, y)

		}
	}
	return count
}

// without visited
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var explore func(x, y int)
	explore = func(x, y int) {
		if x < 0 || x > rows-1 || y < 0 || y > cols-1 || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		explore(x+1, y)
		explore(x, y+1)
		explore(x-1, y)
		explore(x, y-1)
	}
	count := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == '0' {
				continue
			}
			count++
			explore(x, y)

		}
	}
	return count
}
