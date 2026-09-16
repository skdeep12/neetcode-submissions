func largestRectangleArea(h []int) int {
	s := make([]int, 0)
	ans := 0
	for i:=0;i<=len(h);i+=1{
		current := 0
		if i < len(h) {
			current = h[i]
		}
		for len(s) > 0 && h[s[len(s)-1]] > current {
			top := h[s[len(s)-1]]
			s = s[:len(s)-1]
			width := i
			if len(s) > 0{
				width = i-s[len(s)-1]-1
			}
			area := top*width
			ans = max(area,ans)
		}
		s = append(s, i)
	}
	return ans
}
