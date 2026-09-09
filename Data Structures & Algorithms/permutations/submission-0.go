import "slices"
func permute(nums []int) [][]int {
	picked := make(map[int]bool)
	n := len(nums)
	var recurse func([]int)

	finalAns := make([][]int, 0)
	recurse = func(ans []int) {
		if len(ans) == n {
			finalAns = append(finalAns, slices.Clone(ans))
			return
		}
		for i:=0;i<n;i+=1{
			if _, ok:= picked[i]; !ok {
				picked[i] = true
				recurse(append(ans, nums[i]))
				delete(picked, i)
			}
		}
	}
	recurse([]int{})
	return finalAns
}
