func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
	for _, n := range nums{
		if _, ok := m[n]; ok{
			return true
		} else{
			m[n] = true
		}
	}
	return false
}
