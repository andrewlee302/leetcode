func longestPalindromeSubseq(s string) int {
    if len(s) == 0 { return 0 }
    maxL := 1
    dp := make([][]int, len(s))
    for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(s)) }
    for l := 1; l <= len(s); l++ {
        for i := 0; i <= len(s)-l; i++ {
            j := i+l-1
            if l == 1 { dp[i][j] = 1 }
            if l == 2 { if s[i] == s[j] { dp[i][j] = 2 } else { dp[i][j] = 1 } }
            if l > 2 {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
                if s[i] == s[j] { dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2) }
            }
            maxL = max(maxL, dp[i][j])
        }
    }
    return maxL
}

func max(i, j int) int { if i > j { return i } else { return j } }
