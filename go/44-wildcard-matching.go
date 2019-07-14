// DP, like 10. Regular Expression Matching.
func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s)+1)
    for i := 0; i < len(s)+1; i++ { dp[i] = make([]bool, len(p)+1) }
    dp[0][0] = true
    for i := 0; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if i == 0 {
                if p[j-1] == '*' { dp[i][j] = dp[i][j-1] }
            } else if s[i-1] == p[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i-1][j] || dp[i][j-1] // Note!!!
            } else if p[j-1] == '?' {
                dp[i][j] = dp[i-1][j-1]
            }
        }
    }
    return dp[len(s)][len(p)]
}

// Note: optimization 
// before: else if p[j-1] == '*', dp[i][j] = dp[i][j-1] || dp[i-1][j-1] || .. || dp[0][j-1]
// modified: else if p[j-1] == '*', dp[i-1][j] || dp[i][j-1]

// i: 0~len(s), j: 0~len(p)
//
// dp[i][j] 
// if s[i-1] == p[j-1], dp[i][j] = dp[i-1][j-1]
// else if p[j-1] == '*', dp[i-1][j] || dp[i][j-1]
// else if p[j-1] == '?' && i > 0, dp[i][j] = dp[i-1][j-1]
// else dp[i][j] = false
// 
// corner case:
// dp[0][0] = true
// dp[i][0] = false "i>=1"

