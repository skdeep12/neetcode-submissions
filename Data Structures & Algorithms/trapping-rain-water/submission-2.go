func trap(height []int) int {
	s := make([]int, 0)
	ans := 0
	for i:=0;i<len(height);i+=1{
		for len(s) > 0 &&  height[s[len(s)-1]] < height[i] {
			top := s[len(s)-1]
			
			s = s[:len(s)-1]
			if len(s) > 0 {
				left := s[len(s)-1]
				height := min(height[i], height[left]) - height[top]
				width := i-left-1
				ans +=  height*width
			}
			
				
			
			// fmt.Println("inner", i,ans,s)
		}
		s = append(s,i)
		// fmt.Println(i,ans,s)
	}
	
	return ans
}
