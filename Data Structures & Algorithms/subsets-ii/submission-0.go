import "slices"
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	slices.Sort(nums)
	var recurse func(int, []int)
	recurse = func(idx int, a []int) {
		ans = append(ans, slices.Clone(a))
		for i:=idx+1;i<n;i+=1{
			if i>idx+1 && nums[i] == nums[i-1] {
				continue
			}
			recurse(i, append(a, nums[i]))
		}
	}
	recurse(-1, []int{})
	return ans
}
