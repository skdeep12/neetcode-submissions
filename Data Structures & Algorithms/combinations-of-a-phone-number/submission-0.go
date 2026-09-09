func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4' : "ghi",
		'5': "jkl",
		'6' : "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var recurse func(int, string)
	finalAns := make([]string, 0)


	recurse = func(i int, ans string) {
		if i == len(digits) {
			finalAns = append(finalAns, ans)
			return
		}
		keys, _ := m[digits[i]]
		for _,k := range keys {
			recurse(i+1, ans+string(k))
		}
	}

	recurse(0,"")
	return finalAns
}
