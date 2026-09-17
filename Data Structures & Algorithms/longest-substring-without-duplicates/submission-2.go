func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	m := make(map[byte]bool)
	start := 0
	ans := 1
	for i:=0;i<len(s);i+=1{
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = true
		} else {
			for j:=start;j<=i;j+=1{
				if s[j] == s[i]{
					start = j+1
					break
				}
				delete(m,s[j])
			}
		}
		ans = max(ans, i-start+1)
	}
	return ans
}
