package string

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// reverse 倒转字节数组 s将会被修改
func reverse(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		// 交换头尾位置
		tmp := s[i]
		s[i] = s[n-i-1]
		s[n-i-1] = tmp
	}
}

// 三次倒转
func reverseLeftWords2(s string, n int) string {
	bytes := []byte(s)
	reverse(bytes[:n])
	reverse(bytes[n:])
	reverse(bytes)
	return string(bytes)
}
