var ans []int
func spiralOrder(matrix [][]int) []int {
    ans = make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	minrow, maxrow, mincol, maxcol := 0, m-1, 0, n-1
	for i:=0;minrow<=maxrow && mincol<=maxcol;i+=1{
		iter(matrix, minrow, maxrow, mincol, maxcol)
		minrow+=1
		maxrow-=1
		mincol+=1
		maxcol-=1
		// fmt.Println(ans)
	}

	return ans[:m*n]
}



// each iteration
// minrow, maxrow, mincol, maxcol
// 0,m 0,n
// minrow+1, maxrow-1, mincol+1, maxcol-1


func iter(matrix [][]int, minrow, maxrow, mincol, maxcol int) {
	for i:=mincol;i<=maxcol;i+=1{
		ans = append(ans, matrix[minrow][i])
	}
	for j:=minrow+1;j<=maxrow;j+=1{
		ans = append(ans, matrix[j][maxcol])
	}
	for i:=maxcol-1;i>=mincol;i-=1{
		ans = append(ans, matrix[maxrow][i])
	}
	for j:=maxrow-1;j>minrow;j-=1{
		ans = append(ans, matrix[j][mincol])
	}
}
/*


0,0 -> 
0,2 

row, col, maxCol, minCol

maxCol -= 1, minCol+=1

0,0 -> col+=1 
*/