package lc006_easy

// 栈
func isValid(s string) bool {
	stack := []byte{}
	n := len(s)

	match := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := 0; i < n; i++ {
		if match[s[i]] > 0 {
			//栈无元素则无法匹配 ，栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != match[s[i]] {
				return false
			}
			//如果是右括号则匹配出栈
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号则入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
