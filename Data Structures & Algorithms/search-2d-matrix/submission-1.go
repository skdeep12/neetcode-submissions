func searchMatrix(matrix [][]int, target int) bool {
	i := searchRow(matrix, target)
	l := 0
	r := len(matrix[0])-1
	for l < r {
		m := (l+r)/2
		if matrix[i][m] == target {
			return true
		} else if target >  matrix[i][m] {
			l = m+1
		} else {
			r = m-1
		}
	}
	return matrix[i][l] == target
}
func searchRow(matrix [][]int, target int) int {
	l := 0
	r := len(matrix)-1
	for l <r {
		m := (l+r)/2
		if matrix[m][0] == target {
			return m
		} else if target < matrix[m][0] {
			r = m-1
		} else {
			if target <= matrix[m][len(matrix[0])-1] {
				return m
			} else {
				l = m+1
			}
		}
	}
	return l
}