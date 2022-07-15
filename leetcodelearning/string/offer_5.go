package string

// 遍历添加 时间复杂度O(N),空间复杂度O(N)
func replaceSpace(s string) string {
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b = append(b, []byte("%20")...)
		} else {
			b = append(b, s[i])
		}
	}
	return string(b)
}

// 原地修改 时间复杂度O(N),空间复杂度O(1)
func replaceSpace(s string) string {
	b := []byte(s)
	oriLen := len(b)
	// 找出空格个数
	spaceCount := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			spaceCount++
		}
	}
	// 扩充原切片
	tmp := make([]byte, spaceCount*2)
	b = append(b, tmp...)

	l, r := oriLen-1, len(b)-1
	for l <= r {
		if b[l] != ' ' {
			b[r] = b[l]
			l--
			r--
		} else {
			b[r], b[r-1], b[r-2] = '0', '2', '%'
			r -= 3
			l--
		}
	}
	return string(b)
}
