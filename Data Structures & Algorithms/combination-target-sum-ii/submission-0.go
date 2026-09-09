import "slices"
func combinationSum2(candidates []int, target int) [][]int {
	var recurse func(int, []int, int)
	ans := make([][]int, 0)
	
	slices.Sort(candidates)
	

	recurse = func(start int, sol []int, rem int) {
		if rem == 0 {
			a := slices.Clone(sol)
			ans = append(ans, a)
			return
		}
		for i:=start;i<len(candidates);i+=1{
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			if candidates[i] > rem{
				break
			}
			recurse(i+1, append(sol, candidates[i]), rem-candidates[i])
		}
	}
	recurse(0,[]int{},target)
	return ans
}


