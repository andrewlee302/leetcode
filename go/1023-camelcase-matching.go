func camelMatch(queries []string, pattern string) []bool {
    ret := make([]bool, len(queries))
    dp := make([][]bool, 101)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, 101)
    }
    dp[0][0] = true
    for i := 0; i < len(ret); i++ {
        ret[i] = match(queries[i], pattern, dp)
    }
    return ret
}

func match(query string, pattern string, dp [][]bool) bool {
    for i := 1; i < len(query) + 1; i++ {
        if query[i-1] >= 'a' && query[i-1] <= 'z' { dp[i][0] = dp[i-1][0] } else { dp[i][0] = false } // Note
        for j := 1; j < len(pattern) + 1; j++ {
            if query[i-1] != pattern[j-1] {
                if query[i-1] >= 'a' && query[i-1] <= 'z' { dp[i][j] = dp[i-1][j] } else { dp[i][j] = false }
            } else {
                if query[i-1] >= 'a' && query[i-1] <= 'z' { dp[i][j] = dp[i-1][j-1] || dp[i-1][j] } else { dp[i][j] = dp[i-1][j-1] }
            }
        }
    }
    return dp[len(query)][len(pattern)]
}
// dp[xxx][0] = uncertain
// dp[0][xxx] = false
// dp[i][j] = 
// if q[i] != pattern[j] { lowercase dp[i-1][j], uppercase false }
// == { lowercase dp[i-1][j-1] || dp[i-1][j], uppercase dp[i-1]}[j] }
