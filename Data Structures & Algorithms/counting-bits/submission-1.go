func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	} 
	if n == 1{
		return []int{0,1}
	}
	ans := make([]int, n+1)
	ans[0] = 0
	ans[1] = 1
	num := 2
	for i:=2;i<n+1;i+=1{
		if i == num {
			ans[i] = 1
			num *= 2
		} else {
			ans[i] = 1 + ans[i-(num/2)]
		}
	}
	return ans
}
