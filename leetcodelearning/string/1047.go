package string

func removeDuplicates(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && s[i] == stack[len(stack)-1] {
			// 待入栈的字母与栈顶字母相同
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
