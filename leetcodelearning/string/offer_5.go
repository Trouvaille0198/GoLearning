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
func replaceSpace2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	tmp := make([]byte, spaceCount*2)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	// 从后向前填充
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}
