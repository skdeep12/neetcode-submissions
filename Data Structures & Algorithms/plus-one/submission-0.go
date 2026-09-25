func plusOne(digits []int) []int {
	add := 1
    for i:=len(digits)-1;i>=0 && add == 1;i-=1{
		digits[i] += 1
		if digits[i] >= 10 {
			add = 1
			digits[i] -= 10
		} else {
			add = 0
		}
	}
	if add == 1{
		return append([]int{1}, digits...)
	}
	return digits
}
