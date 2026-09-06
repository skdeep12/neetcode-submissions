func jump(nums []int) int {
    n := len(nums)
	if n < 2{
		return 0
	}
	idx := 0
	ans :=0
	currentEnd := 0
	for i:=0;i<n-1;i+=1{
		if idx < i + nums[i] {
			idx = i+nums[i]
		}
		if currentEnd == i{
			ans += 1
			currentEnd = idx
		}
	}
	return ans
}
