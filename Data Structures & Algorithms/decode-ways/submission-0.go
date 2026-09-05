func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
    dp := make([]int, n+1)
	dp[n] = 1
	for i:=n-1;i>=0;i-=1{
		if s[i] != 48 {
			if i<n-1 {
				dp[i] += dp[i+1]
				in,_ := strconv.Atoi(s[i:i+2])
				if in <= 26 {
					dp[i] += dp[i+2]
				}
			} else {
				dp[i] = 1
			}
		} else {
			if i == 0{
				return 0
			}
			in, _ := strconv.Atoi(s[i-1:i+1])
			if in  > 26 {
				return 0
			}
		}
	}
	return dp[0]
} 
