func isAnagram(s string, t string) bool {
	m := make(map[byte][]int)
	if len(s) != len(t) {
		return false
	}
	for idx, _ := range s{
		if val, ok := m[s[idx]]; ok {
			val[0] += 1
		} else {
			m[s[idx]] = []int{1, 0}
		}
		if val, ok := m[t[idx]]; ok {
			val[1] += 1
		} else {
			m[t[idx]] = []int{0,1}
		}
	}
	for _, v := range m {
		if v[0] != v[1] {
			return false
		}
	}
	return true
}
