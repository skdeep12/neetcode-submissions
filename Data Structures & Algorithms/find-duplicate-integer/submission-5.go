func findDuplicate(nums []int) int {
	
	for i := 0;i<len(nums);{

		if nums[i] == i+1 {
            i++
            continue
        }
		correct := nums[i]-1

		if nums[correct] == nums[i] {
			return nums[i]
		}
		nums[i], nums[correct] = nums[correct], nums[i]
	}
	return -1
}
