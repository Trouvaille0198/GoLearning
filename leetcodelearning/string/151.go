package string

func reverse(b []byte) {
	l, r := 0, len(b)-1
	for l <= r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

func reverseWords(s string) string {
	b := []byte(s)
	// 移除多余空格
	slow, fast := 0, 0

	for fast < len(b) {
		// 寻找一个单词的开头
		for fast < len(b) && b[fast] == ' ' {
			fast++
		}
		for fast < len(b) && b[fast] != ' ' {
			b[slow], b[fast] = b[fast], b[slow]
			slow++
			fast++
		}
		slow++
	}
	if b[slow-2] == ' ' {
		// 字符末尾有空格
		b = b[:slow-2]
	} else {
		// 字符末尾无空格
		b = b[:slow-1]
	}

	reverse(b)
	slow, fast = 0, 0
	for fast < len(b) {
		for slow < len(b) && b[slow] == ' ' {
			slow++
		}
		fast = slow
		for fast < len(b) && b[fast] != ' ' {
			fast++
		}
		reverse(b[slow:fast])

		slow = fast
	}
	return string(b)
}
