// dp iterative
func numDecodings(s string) int {
    if len(s) == 0 { return 0 }
    if s[0] == '0' { return 0 }
    dp := make([]int, len(s)+1)
    dp[0], dp[1] = 1, 1
    for i := 2; i < len(dp); i++ {
        if s[i-1] == '0' {
            if s[i-2] == '1' || s[i-2] == '2' { dp[i] = dp[i-2] } else { return 0 }
        } else if s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6' {
            dp[i] = dp[i-2] + dp[i-1]
        } else {
            dp[i] = dp[i-1]
        }
    }
    return dp[len(s)]
}

// dp[0] = 1
// dp[1] = 1 if not '0', otherwise 0
