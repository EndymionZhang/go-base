package main

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			char := stack[len(stack)-1]
			if (c == ')' && char != '(') || (c == ']' && char != '[') || (c == '}' && char != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func isValidV2(s string) bool {
	if len(s) < 2 || len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))
	matchingPair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			char := stack[len(stack)-1]
			if matchingPair[c] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	isValid("([])")
}
