func solveSudoku(board [][]byte)  {
    var rows, cols, blocks [9][10]bool
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            b := (r/3)*3+c/3
            if v := board[r][c]; v != '.' {
                v = v - '0'
                rows[r][v], cols[c][v], blocks[b][v] = true, true, true
            }
        }
    }
    helper(board, 0, 0, rows, cols, blocks)
}

func helper(board [][]byte, r, c int, rows, cols, blocks [9][10]bool) bool {
    if r == 9 && c == 0 { return true }
    if board[r][c] != '.' {
        nextR, nextC := nextRC(r, c)
        return helper(board, nextR, nextC, rows, cols, blocks)
    }
    for v := 1; v <= 9; v++ {
        b := (r/3)*3+c/3
        if !rows[r][v] && !cols[c][v] && !blocks[b][v] {
            board[r][c] = byte(v) + '0'
            nextR, nextC := nextRC(r, c)
            rows[r][v], cols[c][v], blocks[b][v]  = true, true, true
            if helper(board, nextR, nextC, rows, cols, blocks) { return true }
            rows[r][v], cols[c][v], blocks[b][v]  = false, false, false
            board[r][c] = '.'
        }
    }
    return false
}

func nextRC(r, c int) (int, int) {
    if c++; c == 9 { r, c = r+1, 0 }
    return r, c
}
