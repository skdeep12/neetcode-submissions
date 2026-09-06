func canJump(nums []int) bool {
	n := len(nums)
    // dp := make([]int, n)
	idx := 0
	for i:=0;i<n;i+=1{
		if i > idx {
			return false
		}
		idx = max(idx, i+nums[i])
	}
	return idx>=n-1
}
