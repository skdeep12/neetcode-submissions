func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers)-1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum > target {
			end-=1
		} else if sum < target{
			start+=1
		} else {
			return []int{start+1, end+1}
		}
	}
	return []int{}
}
