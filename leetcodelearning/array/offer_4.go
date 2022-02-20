package array

// 行二分
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		low, high := 0, len(nums)-1
		for low <= high {
			mid := (low + high) >> 1
			if nums[mid] > target {
				high = mid - 1
			} else if nums[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// 标志数
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0 // 从左下角开始

	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}
