func solveNQueens(n int) [][]string {
    if n <= 0 { return nil }
    board := make([][]byte, n)
    for row := 0; row < n; row++ {
        board[row] = make([]byte, n)
        for col := 0; col < n; col ++ { board[row][col] = '.' }
    }
    var ret [][]string
    tryPlace(board, 0, &ret)
    return ret
}

func tryPlace(board [][]byte, row int, ret *[][]string) {
    if row >= len(board) { 
        solution := make([]string, len(board))
        for row := 0; row < len(board); row++ { solution[row] = string(board[row]) }
        *ret = append(*ret, solution)
        return
    }
    for col := 0; col < len(board); col++ {
        if isSafe(board, row, col) {
            board[row][col] = 'Q'
            tryPlace(board, row+1, ret)
            board[row][col] = '.'
        }
    }
}

func isSafe(board [][]byte, row, col int) bool {
    // check upper rows conflict
    for i := 0; i < row; i++ { if board[i][col] == 'Q' { return false } }
    // check left diagonal
    for i,j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 { 
        if board[i][j] == 'Q' { return false }
    }
    // check right diagonal
    for i,j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1{ 
        if board[i][j] == 'Q' { return false }
    }
    return true
}
