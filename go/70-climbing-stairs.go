func climbStairs(n int) int {
    if n == 1 { return 1 }
    x, y := 1, 1
    for i := 2; i <= n; i++ {
        x, y = y, x + y
    }
    return y
}
// dp[i] = dp[i-2] + dp[i-1]
// dp[0] = 1, dp[1] = 1
