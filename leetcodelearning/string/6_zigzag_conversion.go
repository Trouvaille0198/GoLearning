package string

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

func convert(s string, numRows int) string {
	step := numRows*2 - 2 // 到达下一层要跳转的步长
	if step <= 0 {
		return s
	}
	result := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		// 遍历每一层
		for j := i; j < len(s); j += step {
			// 控制步长来遍历一层中的每一个元素
			result = append(result, s[j])
			row := j%step + 1 // j所在层数 以1为始
			if row > 1 && row <= numRows-1 && j+(numRows-row)*2 < len(s) {
				// 添加z型拐弯处的字母
				result = append(result, s[j+(numRows-row)*2])
			}
		}
	}
	return string(result)
}
