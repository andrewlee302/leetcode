func divisorGame(N int) bool {
    dp := make([]bool, N+1)
    dp[1] = false
    for i := 2; i <= N; i++ {
        win := false
        for j := 1; j < i; j++ { // find factor
            if i%j == 0 {
                win = win || !dp[i-j]
            }
        }
        dp[i] = win
    }
    return dp[N]
}
