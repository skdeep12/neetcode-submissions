func ladderLength(beginWord string, endWord string, wordList []string) int {
    m := make(map[string][]string, 0)
	wordList = append(wordList, beginWord)
	for idx, w := range wordList {
		if _, ok := m[w];!ok{
			m[w] = make([]string,0)
		} 
		for i:=idx+1;i<len(wordList);i+=1{
			if diff(w,wordList[i]) == 1 {
				m[w] = append(m[w], wordList[i])
				if _, ok := m[wordList[i]]; !ok {
 					m[wordList[i]] = make([]string,0)
				}
				m[wordList[i]] = append(m[wordList[i]], w)
			}
		} 
	}

	q := []string{beginWord}
	visited := make(map[string]bool)
	ans := 0
	// fmt.Println(m)
	for len(q) > 0 {

		l := len(q)
		// fmt.Println(q)
		ans += 1
		for i:=0;i<l;i+=1{
			visited[q[i]] = true
			if q[i] == endWord {
				return ans
			}
			for _, v := range m[q[i]] {
				if !visited[v] {
					q = append(q,v)
				}
			}
		}
		q = q[l:]
	}
	return 0
}

func diff(s1, s2 string) int {
	d := 0
	for i:=0;i<len(s1);i+=1{
		if s1[i] != s2[i] {
			d+=1
		}
	}
	return d
}
