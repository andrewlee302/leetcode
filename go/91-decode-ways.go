// DP iterative.
func numDecodings(s string) int {
    if len(s) == 0 { return 0 }
    dp := make([]int, len(s))
    if s[0] != '0' { dp[0] = 1 }
    for i := 1; i < len(s); i++ {
        if s[i] != '0' {
            dp[i] += dp[i-1]
        }
        if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '0' && s[i] <= '6' {
            if i == 1 { dp[i] += 1 } else { dp[i] += dp[i-2] }
        }
    }
    return dp[len(s)-1]
}
