package array

// Given n non-negative integers a1, a2, …, an , where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
// Find two lines, which together with x-axis forms a container, such that the container contains the most water.
// Note: You may not slant the container and n is at least 2.

// 对撞指针
func maxArea(height []int) int {
	start, end := 0, len(height)
	curResult, maxResult := 0, 0
	for start < end {
		curWidth := end - start
		curHeight := 0
		// 因为高度取矮者 矮的一边的指针就可以往中间移动 因为反过来的移法面积必定更小
		if height[start] < height[end] {
			curHeight = height[start]
			start++
		} else {
			curHeight = height[end]
			end--
		}

		curResult = curWidth * curHeight
		if curResult > maxResult {
			maxResult = curResult
		}
	}
	return maxResult
}
