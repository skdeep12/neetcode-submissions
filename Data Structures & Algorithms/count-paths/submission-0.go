func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
	for i:=0;i<m;i+=1{
		dp[i] = make([]int,n)
		dp[i][0] = 1
		for j:=1;j<n;j+=1{
			if i == 0{
				dp[i][j] = 1
			} else  {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
