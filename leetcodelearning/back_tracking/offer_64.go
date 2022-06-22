package back_tracking

// 逻辑短路
func sumNums(n int) int {
	res := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1) // 不满足n>1时 递归就会退出
	}
	sum(n)
	return res
}
