func islandsAndTreasure(grid [][]int) {
	m := len(grid)
	if m==0{
		return
	}
	n := len(grid[0])
	q := make([][]int, 0)
	for i:=0;i<m;i+=1{
		for j:=0;j<n;j+=1{
			if grid[i][j] == 0 {
				q = append(q, []int{i,j, 0})
			}
		}
	}
	for len(q) > 0{
		l := len(q)
		for i:=0;i<l;i+=1{
			current := q[i]
			for _, val := range [][]int{[]int{1,0},[]int{-1,0},[]int{0,1},[]int{0,-1}}{
				x := current[0] + val[0]
				y := current[1] + val[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == -1 {
					continue
				}
				if grid[x][y] > current[2] + 1 {
					grid[x][y] = current[2] + 1
					q = append(q, []int{x,y,current[2]+1})
				}
			}
		}
		q = q[l:]
	}
}
