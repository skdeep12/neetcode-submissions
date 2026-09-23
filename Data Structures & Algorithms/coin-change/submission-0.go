func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i:=1;i<amount+1;i+=1{
		t := -1
		for _, c := range coins{
			if i == c{
				t = 1
			} else if i-c > 0 && dp[i-c] != -1{
				if t == -1{
					t = 1+dp[i-c]
				} else {
					t = min(t,1+dp[i-c])
				}
			}
		}
		dp[i] = t
	}
	return dp[amount]
}
