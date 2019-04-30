func maxUncrossedLines(A []int, B []int) int {
    dp := make([][]int, len(A)+1)
    for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(B)+1) }
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if A[i-1] == B[j-1] {
               dp[i][j] = 1 + dp[i-1][j-1]
            } else {
               if dp[i-1][j] > dp[i][j-1] { 
                    dp[i][j] = dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-1]
                } 
            }
        }
    }
    return dp[len(A)][len(B)]
}
