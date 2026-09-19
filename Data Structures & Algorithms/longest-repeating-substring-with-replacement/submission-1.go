func characterReplacement(s string, k int) int {
	// largest string which has, min heap and sum of the rest is 
	// at most k dif
	// AAABABB
	left := 0
	right := 0
	f := make([]int, 26)
	ans := 0
	for right < len(s) {
		f[s[right]-'A']+=1
		for !isValid(f,k){
			// fmt.Println(f)
			f[s[left]-'A']-=1
			left+=1
		}
		ans = max(ans, right-left+1)
		right+=1
		
	}
	return ans
}

func isValid(f []int, k int) bool {
	m := 0
	l := 0
	for _, r := range f{
		m = max(m,r)
		l+=r
	}
	return (l-m) <= k
} 
