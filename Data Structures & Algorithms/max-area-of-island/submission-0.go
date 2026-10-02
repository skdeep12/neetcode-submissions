func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
    if m == 0{
        return 0
    }
    n := len(grid[0])
    visited := make([][]int, m)

    for i:=0;i<m;i+=1{
        visited[i] = make([]int, n)
    }
    ans := 0
    for i:=0;i<m;i+=1{
        for j:=0;j<n;j+=1{
            if grid[i][j] == 1 && visited[i][j] == 0 {
                ans = max(ans, dfs(grid, i, j, visited))
            }
        }
    }
    return ans
}

func dfs(grid [][]int, i, j int, visited [][]int) int {
    m := len(grid)
    n := len(grid[0])
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0  || visited[i][j] == 1{
        return 0
    }
    visited[i][j] = 1
    a := 1
    a += dfs(grid,i+1,j,visited)
    a += dfs(grid,i-1,j,visited)
    a += dfs(grid,i,j+1,visited)
    a += dfs(grid,i,j-1,visited)
    return a
}
