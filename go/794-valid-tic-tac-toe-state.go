func validTicTacToe(board []string) bool {
    cntX, cntO := 0, 0
    for _, row := range board {
        for i := 0; i < len(row); i++ {
            if row[i] == 'X' { cntX++ }
            if row[i] == 'O' { cntO++ }
        }
    }
    winX, winO := check(board, 'X'), check(board, 'O')
    if cntX != cntO && cntX != cntO + 1 || winX && winO || winX && cntX == cntO || winO && cntX > cntO { return false }
    return true
}

func check(board []string, p byte) bool {
    for i := 0; i < 3; i++ {
        if row := board[i]; row[0] == p && row[1] == p && row[2] == p { return true }
    }
    for i := 0; i < 3; i++ {
        if board[0][i] == p && board[1][i] == p && board[2][i] == p { return true }
    }
    if board[0][0] == p && board[1][1] == p && board[2][2] == p { return true }
    if board[0][2] == p && board[1][1] == p && board[2][0] == p { return true }
    return false
}
