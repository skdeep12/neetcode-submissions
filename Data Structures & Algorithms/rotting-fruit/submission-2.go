func orangesRotting(grid [][]int) int {
    q := make([][]int, 0)
	wasOne := false
	for i:=0;i<len(grid);i+=1{
		for j:=0;j<len(grid[0]);j+=1{
			if grid[i][j] == 2{
				q = append(q, []int{i,j})
			} else if grid[i][j] == 1{
				wasOne = true
			}
		}
	}
	if len(q) == 0{
		if wasOne{
			return -1
		} else {
			return 0
		}
	}
	ans := 0
	iterator := [][]int{[]int{1,0}, []int{0,1}, []int{-1,0}, []int{0,-1}}
	for len(q) > 0 {
		ans += 1
		l := len(q)
		for i:=0;i<l;i+=1{
			for j:=0;j<4;j+=1{
				x := q[i][0] + iterator[j][0]
				y := q[i][1] + iterator[j][1]
				if x >= 0 && x<len(grid) && y >=0 && y< len(grid[0]) && grid[x][y] == 1{					grid[x][y] = 2
					q = append(q, []int{x,y})
				}
			}
		}
		q = q[l:]
		
	}
	for i:=0;i<len(grid);i+=1{
		for j:=0;j<len(grid[0]);j+=1{
			if grid[i][j] == 1{
				return -1
			}
		}
	}
	return ans-1
}
