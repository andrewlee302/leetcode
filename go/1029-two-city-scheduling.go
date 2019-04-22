func twoCitySchedCost(costs [][]int) int {
    N := len(costs)/2
    dp := make([][]int, 2*N+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, N+1)
    }
    dp[0][0] = 0
    for i := 1; i < len(dp); i++ {
        for j := 0; j <= i && j <= N; j++ {
            if i != j {
                dp[i][j] = dp[i-1][j] + costs[i-1][1] // B
                if j > 0 {
                    temp := dp[i-1][j-1] + costs[i-1][0] // A
                    if temp < dp[i][j] { dp[i][j] = temp }
                }
            } else {
                dp[i][j] = dp[i-1][j-1] + costs[i-1][0] // A
            }
        }
    }
    fmt.Println(dp)
    return dp[2*N][N]
}
