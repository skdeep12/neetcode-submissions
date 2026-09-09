
func solveNQueens(n int) [][]string {
	ans := make([][]rune, n)
	for i:=0;i<n;i+=1{
		ans[i] = getStr(n)
	}
	finalAns := make([][]string, 0)
	var recurse func(int)

	recurse = func(i int){
		if i == n{
			local := make([]string, 0)
			for i:=0;i<n;i+=1{
				local = append(local, string(ans[i]))
			}
			finalAns = append(finalAns, local)
			return
		}

		for j:=0;j<n;j+=1{
			if checkPlacement(ans, i, j){
				ans[i][j] = 'Q'
				recurse(i+1)
			}
			ans[i][j] = '.'
			
		}
	}
	recurse(0)
	return finalAns
}

func checkPlacement(board [][]rune, i, j int) bool{
	for k:=0;k<i;k+=1{
		if board[k][j] == 'Q' {
			return false
		}
	}
	m := i
	n := j
	for m >= 0 && n < len(board) {
		if board[m][n] == 'Q' {
			return false
		}
		m-=1
		n+=1
	}
	m = i
	n = j
	for m >= 0 && n >= 0 {
		if board[m][n] == 'Q' {
			return false
		}
		m-=1
		n-=1
	}
	// check diagonals
	return true
}

func getStr(n int) []rune{
	ans := make([]rune, 0)
	for i:=0;i<n;i+=1{
		ans = append(ans, '.')
	}
	return ans
}

