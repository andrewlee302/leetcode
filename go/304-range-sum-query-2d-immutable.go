type NumMatrix struct {
    dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    dp := make([][]int, len(matrix)+1)
    if len(matrix) == 0 { return NumMatrix{dp:dp} }
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(matrix[0])+1) // default 0
    }
     for i := 1; i < len(dp); i++ {
        rowSum := 0
        for j := 1; j < len(dp[0]); j++ {
            rowSum += matrix[i-1][j-1]
            dp[i][j] = dp[i-1][j] + rowSum
        }
    }
    return NumMatrix{dp:dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    if !(row1 >= 0 && col1 >= 0 && row2 < len(this.dp) && col2 < len(this.dp[0])) {
        return -1
    }
    return this.dp[row2+1][col2+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1] + this.dp[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
