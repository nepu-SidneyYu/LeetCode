package InterviewQuestions

func isValid1(s string) bool {
	stack := make([]byte, 0)
	ss := []byte(s)
	for _, a := range ss {
		if a == ')' || a == ']' || a == '}' {
			n := len(stack)
			if len(stack) == 0 {
				return false
			} else if stack[n-1] == '(' && a == ')' {
				stack = stack[:n-1]
			} else if stack[n-1] == '[' && a == ']' {
				stack = stack[:n-1]
			} else if stack[n-1] == '{' && a == '}' {
				stack = stack[:n-1]
			}
		} else {
			stack = append(stack, a)
		}

	}
	return len(stack) == 0
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	ss := []byte(s)
	pair := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}
	for _, a := range ss {
		if pair[a] > 0 {
			if len(stack) == 0 || pair[a] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, a)
		}
	}
	return len(stack) == 0
}
