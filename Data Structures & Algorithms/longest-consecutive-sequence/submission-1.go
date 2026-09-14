func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, t := range nums{
		m[t] = true
	}
	ans := 0
	for _, t := range nums{
		if _, ok := m[t-1]; !ok {
			a := check(nums, m, t)
			ans = max(a,ans)
		}
	}
	return ans
}

func check(nums []int, m map[int]bool, start int) int {
	ans := 1 
	for true {
		if _, ok := m[start+1]; !ok {
			return ans
		}
		start+=1
		ans+=1
	}
	return 0
}
