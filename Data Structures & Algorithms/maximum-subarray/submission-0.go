func maxSubArray(nums []int) int {
    ans := math.MinInt
	sum := math.MinInt
	for _, n := range nums {
		if sum < 0 {
			sum = n
		} else {
			sum += n
		}
		ans = max(sum ,ans)
	}
	return ans
}
