type NumMatrix struct {
    dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 { // Note.
        return NumMatrix{nil}
    }
    dp := make([][]int, len(matrix) + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(matrix[0]) + 1)
    }
    for i := 1; i < len(dp); i++ {
        sum := 0
        for j := 1; j < len(dp[0]); j++ {
            sum += matrix[i-1][j-1]
            dp[i][j] = dp[i-1][j] + sum
        }
    }
    return NumMatrix{dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    dp := this.dp
    return dp[row2+1][col2+1] - dp[row1][col2+1] - dp[row2+1][col1] + dp[row1][col1]
}

// when there are no dummy slots in dp.
// row1, col1, row2, col2:
// dp[row2][col2] - dp[row1-1][col2] - dp[row2][col1-1] + dp[row1-1][col1-1]
// dp[i][j] = dp[i-1][j] + sum(matrix[i][0..j])
// dp[i][j+1] = dp[i-1][j+1] + sum(matrix[i][0..j+1])

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
