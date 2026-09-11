func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	start := 0
	end := len(s)-1
	for start < end{
		for start < end && shouldSkip(s[start]){
			start+=1
		}
		for start < end && shouldSkip(s[end] ){
			end-=1
		}
		if start >= end {
			break
		}
		
		if strings.ToLower(s[start:start+1]) != strings.ToLower(s[end:end+1]) {
			return false
		}
		start+=1
		end-=1
	}
	return true
}

func shouldSkip(a byte) bool {
    if (a >= '0' && a <= '9') ||
        (a >= 'a' && a <= 'z') ||
        (a >= 'A' && a <= 'Z') {
        return false
    }
    return true
}