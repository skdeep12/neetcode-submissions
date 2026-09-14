func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	ans := 0
	for l < r {
		area := (r-l) * min(heights[l], heights[r])
		if heights[l] < heights[r] {
			l+=1
		} else {
			r-=1
		}
		ans = max(ans, area)
	}
	return ans
}
