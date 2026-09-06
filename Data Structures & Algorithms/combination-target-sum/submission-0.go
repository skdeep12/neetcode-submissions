
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
		ans_2 := make([]int, len(ans))
		copy(ans_2, ans)
		comb = append(comb, ans_2)
		return
	}
	if idx == len(nums) {
		return
	}
	
	recurse(nums, idx+1, sum, ans)
	recurse(nums, idx, sum-nums[idx], append(ans, nums[idx]))
}
