func lengthOfLIS(nums []int) int {
	n := len(nums)
    if n == 0{
		return 0
	}
	dp := make([]int, n)
		
	for i:=0;i<n;i+=1{
		dp[i] = 1
		for j:=i-1;j>=0;j-=1{
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}			
		}
	}
	return arrMax(dp)

}
func arrMax(a []int) int {
	ans := a[0]
	for i:=1;i<len(a);i+=1{
		ans = max(ans,a[i])
	}
	return ans
}