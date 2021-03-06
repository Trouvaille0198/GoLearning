package stdpkglearning

// 手动实现的Reader，过滤输入字符串的非字母部分

type AlphaReader struct {
	src string // 待读取的字符
	cur int    // 当前读到的位置
}

func NewAlphaReader(src string) *AlphaReader {
	return &AlphaReader{src: src}
}

func getAlpha(r byte) byte {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return r
	}
	return 0
}

func (a *AlphaReader) Read(p []byte) (int, error) {
	bound := 0 // 缓冲区大小
	remainLen := len(a.src) - a.cur
	if remainLen >= len(p) {
		// 填满缓冲区
		bound = len(p)
	} else {
		// 剩余字符小于缓冲区大小
		bound = remainLen
	}

	buffer := make([]byte, bound)
	for i := 0; i < bound; i++ {
		char := getAlpha(a.src[a.cur])
		if char != 0 {
			buffer[i] = char
		}
		a.cur++
	}
	copy(p, buffer)
	return bound, nil
}
