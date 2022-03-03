package others

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	res := 1.0
	var b int // 二进制n的某一位 0或1
	for n != 0 {
		b = n & 1 // 获取二进制n最后一位
		if b != 0 {
			res *= x
		}
		n = n >> 1 // n右移一位
		x = x * x
	}
	return res
}
