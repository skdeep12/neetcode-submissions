func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	r := 1
	left[0] = 1
	for i:=1;i<n;i+=1{
		left[i] = left[i-1] * nums[i-1]
	}
	for i:=n-2;i>=0;i-=1{
		r = r * nums[i+1]
		left[i]  = left[i] * r
	}
	return left
}
