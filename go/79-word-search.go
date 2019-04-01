func exist(board [][]byte, word string) bool {
    if len(word) == 0 { return true }
    ret := false
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            ret = ret || DFS(board, word, i, j, 0)
        }
    }
    return ret
}

func DFS(board [][]byte, word string, i, j, k int) bool {
    if k >= len(word) { return true }
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) { return false }
    if board[i][j] != word[k] { return false }
    save := board[i][j]
    board[i][j] = '#'
    ret := DFS(board, word, i, j+1, k+1) || DFS(board, word, i, j-1, k+1) || 
        DFS(board, word, i+1, j, k+1) || DFS(board, word, i-1, j, k+1) 
    board[i][j] = save
    return ret
}
