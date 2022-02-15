package array

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	lowCol, highCol := 0, len(matrix[0])-1
	lowRow, highRow := 0, len(matrix)-1
	for lowCol <= highCol && lowRow <= highRow {
		midCol := (lowCol + highCol) >> 1
		midRow := (lowRow + highRow) >> 1
		if matrix[midRow][midCol] > target {
			highCol = midCol - 1
			highRow = midRow - 1
		} else if matrix[midRow][midCol] < target {
			lowCol = midCol + 1
			lowRow = midRow + 1
		} else {
			return true
		}

	}
	return false
}
