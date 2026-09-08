func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
    parent := make([]int, n+1)
	for i:=0;i<n+1;i+=1{
		parent[i] = i
	}
	for _, e := range edges {
		if union(&parent, e[0], e[1]) {
			return e
		}
	}
	return []int{}
}

func union(parent *[]int, u, v int) bool {
	up := find(*parent, u)
	vp := find(*parent, v)
	if up == vp{
		return true
	}
	(*parent)[vp] = up
	return false
}

func find(parent []int, v int) int{
	for v != parent[v] {
		v = parent[v]
	}
	return v
}
