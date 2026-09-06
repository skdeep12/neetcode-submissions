func isValid(s string) bool {
    stack := make([]byte, 0)
	for i:=0;i<len(s);i+=1{
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			n := len(stack)
			if n == 0 {
				return false
			}
			top := stack[n-1]
			if (s[i] == ')' && top != '(') ||   (s[i] == '}' && top != '{') ||  (s[i] == ']' && top != '[') {
				return false
			}
			stack = stack[:n-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
