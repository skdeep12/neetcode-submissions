import "slices"
func subsets(nums []int) [][]int {

	var recurse func(int, []int)
	n := len(nums)
	ans := make([][]int, 0)
	recurse = func(idx int, a []int){
		ans = append(ans,slices.Clone(a))

		for i:=idx;i<n;i+=1{
			recurse(i+1, append(a,nums[i]))
		}
	}
	recurse(0,[]int{})
	return ans
}

