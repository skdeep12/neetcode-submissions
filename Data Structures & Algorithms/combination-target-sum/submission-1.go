
import "slices"
var comb [][]int
func combinationSum(nums []int, target int) [][]int {
    comb = make([][]int, 0)
	recurse(nums, 0, target, []int{})
	return comb
}

func recurse(nums []int, idx int, sum int, ans []int) {
	// fmt.Println(idx, sum ,ans)
	if sum < 0{
		return
	} else if sum == 0{
		comb = append(comb, slices.Clone(ans))
		return
	}
	if idx == len(nums) {
		return
	}
	for i:=idx;i<len(nums);i+=1{
		recurse(nums, i, sum-nums[i], append(ans, nums[i]))
	}
}
