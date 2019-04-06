func isMatch(s string, p string) bool {
    dp := make([][]bool, len(p)+1)
    for i := 0; i < len(p)+1; i++ {
        dp[i] = make([]bool, len(s)+1)
    }
    dp[0][0] = true // dp[0][0..len(s)] = false, dp[1][0] = false, dp[0..len(p)][0] uncertain
    for i := 2; i < len(p) + 1; i++ {
        // p[0] mustn't be *
        if p[i-1] == '*' { dp[i][0] = dp[i-2][0] }
    }
    for i := 1; i < len(p) + 1; i++ {
        for j := 1; j < len(s) + 1; j++ {
            if p[i-1] >= 'a' && p[i-1] <= 'z' {
                dp[i][j] = p[i-1] == s[j-1] && dp[i-1][j-1]
            } else if i >= 2 && p[i-1] == '*' && p[i-2] != '.' {
                flag := dp[i-2][j]
                for k := 1; k <= j; k++ {
                    if p[i-2] != s[j-k] { break }
                    flag = flag || dp[i-2][j-k]
                }
                dp[i][j] = flag
            } else if i >= 2 && p[i-1] == '*' && p[i-2] == '.' {
                flag := dp[i-2][j]
                for k := 1; k <= j; k++ { flag = flag || dp[i-2][j-k] }
                dp[i][j] = flag
            } else {
                dp[i][j] = dp[i-1][j-1]
            }
        }
    }
    return dp[len(p)][len(s)]
}
