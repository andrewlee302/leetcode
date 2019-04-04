func countBattleships(board [][]byte) int {
    count := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == '.' { continue }
            if i >= 1 && board[i-1][j] == 'X' { continue }
            if j >= 1 && board[i][j-1] == 'X' { continue }
            count++
        }
    }
    return count
}
