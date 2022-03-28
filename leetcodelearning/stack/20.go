package stack

import "container/list"

func isValid(s string) bool {
	quoteMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := list.New()
	stack.PushBack(byte('?')) // 防止第一个括号就是右括号而造成空栈的情况
	for _, quote := range []byte(s) {
		switch quote {
		case ')', ']', '}':
			// 出栈
			popQuote := stack.Back().Value.(byte)
			if quote != quoteMap[popQuote] {
				return false
			} else {
				stack.Remove(stack.Back())
			}
		default:
			stack.PushBack(quote)
		}
	}
	return stack.Len() == 1
}

func isValid(s string) bool {
	quoteMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0)
	for _, quote := range []byte(s) {
		switch quote {
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if quote != quoteMap[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, quote)
		}
	}
	return len(stack) == 0
}
