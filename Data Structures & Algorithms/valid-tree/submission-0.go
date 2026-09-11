func validTree(n int, edges [][]int) bool {
    s := make([]int, 0)
    if len(edges) != n-1 {
        return false
    }
    s = append(s, 0)
    g := make([][]int, n)
    for i := 0;i<n;i+=1{
        g[i] = make([]int, 0)
    }
    for _, e := range edges {
        u,v := e[0], e[1]
        g[u] = append(g[u],v)
        g[v] = append(g[v], u)
    }
    visited := make([]bool, n)

    
    for len(s) > 0 {
        f := s[0]
        s = s[1:]
        visited[f] = true
        for _, v := range g[f]{
            if !visited[v] {
                s = append(s, v)
            }
        }
    }
    for i:=0;i<n;i+=1{
        if !visited[i]{
            return false
        }
    }
    return true
}
