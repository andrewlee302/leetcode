func isValidPalindrome(s string, k int) bool {
    dp := make([][]int, len(s))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(s))
        for j := i + 1; j < len(s); j++ {
            dp[i][j] = len(s)
        }
    }
    for l := 2; l <= len(s); l++ {
        for i := 0; i <= len(s) - l; i++ {
            j := i + l - 1
            if l == 2 {
                if s[i] == s[j] { dp[i][j] = 0 } else { dp[i][j] = 1 }
            } else {
                if s[i] == s[j] {
                    dp[i][j] = dp[i+1][j-1]
                } else {
                    dp[i][j] = min(dp[i+1][j-1]+2, min(dp[i+1][j]+1, dp[i][j-1]+1))
                }
            }
        }
    }
    return dp[0][len(s)-1] <= k
}

func min(a, b int) int { if a < b { return a } else { return b } }
