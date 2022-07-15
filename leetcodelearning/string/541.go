package string

func reverse(b []byte, l, r int) []byte {
	for l <= r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return b
}

func reverseStr(s string, k int) string {
	b := []byte(s)
	curLen := 0
	for curLen < len(s) {
		if curLen+2*k > len(s) {
			if len(s)-curLen > k {
				b = reverse(b, curLen, curLen+k-1)
			} else {
				b = reverse(b, curLen, len(s)-1)
			}
			return string(b)
		} else {
			b = reverse(b, curLen, curLen+k-1)
			curLen += k * 2
		}
	}
	return string(b)
}
