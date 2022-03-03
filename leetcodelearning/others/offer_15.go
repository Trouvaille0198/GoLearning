package others

func hammingWeight(num uint32) (count int) {
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return
}

func hammingWeight2(num uint32) (count int) {
	for num > 0 {
		count++
		num &= num - 1 // 将num最低位的1变成0
	}
	return
}
