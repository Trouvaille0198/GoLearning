package back_tracking

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	if len(word) > rows*cols {
		return false
	}

	var dfs func(layer, x, y int, oriDirection string) bool
	dfs = func(layer, x, y int, oriDirection string) bool {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return false
		}	
		if word[layer] != board[x][y] {
			return false
		}
		if layer == len(word)-1 {
			return true
		}

		board[x][y] = ' '
		upFlag := dfs(layer+1, x-1, y, "down")
		downFlag := dfs(layer+1, x+1, y, "up")
		leftFlag := dfs(layer+1, x, y-1, "right")
		rightFlag := dfs(layer+1, x, y+1, "left")
		board[x][y] = word[layer]

		if upFlag || downFlag || leftFlag || rightFlag {
			return true
		}
		return false

	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			flag := dfs(0, x, y, "")
			if flag {
				return true
			}
		}
	}
	return false
}
