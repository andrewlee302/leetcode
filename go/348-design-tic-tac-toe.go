type TicTacToe struct {
    rows, cols []int
    diagonal, antidiagonal int
    size int
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    return TicTacToe{rows:make([]int, n), cols:make([]int, n), diagonal:0, antidiagonal:0, size:n}
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    add := 1
    if player == 2 { add = -1 }
    this.rows[row] += add
    this.cols[col] += add
    if row == col { this.diagonal += add }
    if row + col == this.size - 1 { this.antidiagonal += add }
    target := add*(this.size)
    if this.rows[row] == target || this.cols[col] == target ||
        this.diagonal == target || this.antidiagonal == target {
        return player
    }
    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
