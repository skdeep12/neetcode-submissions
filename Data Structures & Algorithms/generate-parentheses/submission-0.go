func generateParenthesis(n int) []string {
	var recurse func(int, int)
	ans := ""
	finalAns := make([]string, 0)
	recurse = func(open, clos int ){
		if open == 0 && clos == 0 {
			finalAns=append(finalAns, ans)
			return
		}
		if clos < open || clos == -1 || open == -1 {
			return
		}
		ans += "("
		recurse(open-1, clos)	
		ans = ans[:len(ans)-1]
		ans += ")"
		recurse(open, clos-1)
		ans = ans[:len(ans)-1]
	}
	recurse(n,n)
	return finalAns
}
