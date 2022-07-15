package array

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) (res int) {
	// 用dp计算前后最大高度
	lDp, rDp := make([]int, len(height)), make([]int, len(height))
	lDp[0], rDp[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		lDp[i] = max(lDp[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rDp[i] = max(rDp[i+1], height[i])
	}

	for i := 1; i < len(height)-1; i++ {
		// 左右两个柱子不接雨水
		lMax := lDp[i]
		rMax := rDp[i]
		res += min(lMax, rMax) - height[i]
	}
	return
}

func trap(height []int) (res int) {
	stack := []int{0} // 栈顶到栈底从小到大 存放下标

	for i := 1; i < len(height); i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			// 更新栈顶
			stack[len(stack)-1] = i
		} else {
			for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
				midIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1] // 出栈
				if len(stack) == 0 {
					break
				}
				h := min(height[i], height[stack[len(stack)-1]]) - height[midIndex]
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
			stack = append(stack, i)
		}
	}
	return
}
