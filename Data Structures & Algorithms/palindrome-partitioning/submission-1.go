func partition(s string) [][]string {
	var recurse func(int)[][]string
	recurse = func(start int) [][]string{
		ans := make([][]string, 0)
		for i:=start+1;i<=len(s);i+=1{
			if isPalin(s[start:i]){
				forward_ans := recurse(i)
				for _, f := range forward_ans {
					f = append([]string{s[start:i]}, f...)
					ans = append(ans, f)
				}
				if len(forward_ans) == 0{
					ans = append(ans, []string{s[start:i]})
				}
			}
		}
		return ans
	}
	return recurse(0)
}	

func isPalin(s string) bool {
	if len(s) == 1{
		return true
	}
	i := 0
	j := len(s) - 1
	for i<j {
		if s[i] == s[j]{
			i+=1
			j-=1
		}else{
			break
		}
	}
	return i>=j
}
