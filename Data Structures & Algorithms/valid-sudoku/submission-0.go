func isValidSudoku(board [][]byte) bool {
    mem := make(map[byte]bool)
    var boxes []map[byte]bool
    for i:=0;i<9;i+=1{
        mem = make(map[byte]bool)
        if i%3 == 0{
            boxes = make([]map[byte]bool,3)
            for k:=0;k<3;k+=1{
                boxes[k] = make(map[byte]bool)
            }
        }
        for j:=0;j<9;j+=1{
            if board[i][j] != '.' {
                if _, ok := mem[board[i][j]]; ok {
                    return false
                } else{
                    mem[board[i][j]] = true
                }
                idx := int(j/3)
                if _, ok := boxes[idx][board[i][j]]; ok {
                    return false
                } else{
                    boxes[idx][board[i][j]] = true
                }
            }
        }
    }
    for j:=0;j<9;j+=1{
        mem = make(map[byte]bool)
        for i:=0;i<9;i+=1{
            if board[i][j] != '.' {
                if _, ok := mem[board[i][j]]; ok {
                    return false
                } else{
                    mem[board[i][j]] = true
                }
            }
        }
    }
    return true
}