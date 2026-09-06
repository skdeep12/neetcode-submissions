func dailyTemperatures(t []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(t))
	for i:=0;i<len(t);i+=1{
		if len(stack) == 0{
			stack = append(stack, i)
		} else{ 
			if t[i] < t[stack[len(stack)-1]] {
				stack = append(stack, i)
			} else{
				
				for len(stack) > 0 && t[i] > t[stack[len(stack)-1]] {
					top := stack[len(stack)-1]
					diff := i-top
					ans[top] = diff
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, i)
			}
		}
	}
	return ans
}
