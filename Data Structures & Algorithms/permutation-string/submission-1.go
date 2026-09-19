func checkInclusion(s1 string, s2 string) bool {
	f := make([]int, 26)
	charCount := 0
	for _, s := range s1{
		if f[s-'a'] == 0{
			charCount+=1
		}
		f[s-'a']+=1
	}
	fs := make([]int, 26)
	l:=0
	for idx, s := range s2 {
		right := s-'a'
		fs[right] += 1
		if fs[right] == f[right] && f[right] !=0{
			charCount-=1
		}
		for l <= idx && fs[right] > f[right]{
			left := s2[l]-'a'
			fs[left]-=1
			if fs[left] == f[left]-1 {
				charCount += 1
				// fmt.Println("incremented")
			}
			l+=1
		}
		
		
		// fmt.Println(charCount, l, idx, fs[right], f[right])
		if charCount == 0{
			return true
		}
	}
	return false
}
