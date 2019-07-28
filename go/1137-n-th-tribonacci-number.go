func tribonacci(n int) int {
    if n == 0 { return 0 } else if n == 1 || n == 2 { return 1 }
    dp := make([]int, n+1)
    dp[0], dp[1], dp[2] = 0, 1, 1
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
    }
    return dp[n]
}
