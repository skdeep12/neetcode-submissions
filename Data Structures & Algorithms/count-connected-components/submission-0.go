func countComponents(n int, edges [][]int) int {
    g := make([][]int, n)
	for i:=0;i<n;i+=1{
		g[i] = make([]int, 0)
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	visited := make([]int, n)
	ans := 0
	for i:=0;i<n;i+=1{
		if visited[i] == 0{
			ans += 1
			dfs(g,i,visited)
		}
	}
	return ans
}

func dfs(g [][]int, v int, visited []int) {
	visited[v] = 1
	for i:=0;i<len(g[v]);i+=1{
		if visited[g[v][i]] == 0 {
			dfs(g,g[v][i], visited)
		}
	}
}
