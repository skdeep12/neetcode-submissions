func exist(board [][]byte, word string) bool {
	visited := make([][]int, len(board))
	for i:=0;i<len(board);i+=1{
		visited[i] = make([]int, len(board[0]))
	}
	m := len(board)
	n := len(board[0])
	var dfs func(int,int,int)bool
	dfs = func(i,j int, idx int) bool {
		if idx == len(word) {
			return true
		}
		if i >= m || i<0 || j<0 || j>= n {
			return false
		}
		if board[i][j] != word[idx] || visited[i][j] == 1{
			return false
		}
		visited[i][j] = 1
		
		if dfs(i+1,j, idx+1){
			return true
		}
		if dfs(i-1,j, idx+1){
			return true
		}
		if dfs(i,j+1, idx+1){
			return true
		}
		if dfs(i,j-1, idx+1){
			return true
		}
		visited[i][j] = 0
		return false
	}
	for i:=0;i<len(board);i+=1{
		for j:=0;j<len(board[0]);j+=1{
			if board[i][j] == word[0] {
				if dfs(i,j,0) {
					return true
				}
			}
		}
	}
	return false
}
