package others

func add(a int, b int) int {
	// 进位
	var carry int
	for b != 0 {
		// 进位
		carry = (a & b) << 1
		// 不加进位
		a ^= b
		// 加进位
		b = carry
	}

	return a
}
