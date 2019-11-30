func maxVacationDays(flights [][]int, days [][]int) int {
    N, K := len(flights), len(days[0])
    dp := make([][]int, K)
    for i := 0; i < K; i++ {
        dp[i] = make([]int, N)
        for j := 0; j < N; j++ { dp[i][j] = -1 }
    }
    for i := 0; i < N; i++ {
        if i == 0 || flights[0][i] == 1 {
            dp[0][i] = days[i][0]
        }
    }
    ret := 0
    for i := 1; i < K; i++ {
        for p1 := 0; p1 < N; p1++ {
            for p2 := 0; p2 < N; p2++ {
                if (p1 == p2 || flights[p2][p1] == 1) && dp[i-1][p2] != -1 {
                    dp[i][p1] = max(dp[i][p1], dp[i-1][p2]+days[p1][i])
                }
            }
            ret = max(ret, dp[i][p1])
        }
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }

// dp[k][i], stay in i-th place in k-th week. (0 <= i < N, 0 <= k <= K)
