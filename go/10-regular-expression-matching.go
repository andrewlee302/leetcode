func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s)+1)
    for i := 0; i < len(dp); i++ { dp[i] = make([]bool, len(p)+1) }
    dp[0][0] = true
    for j := 2; j < len(dp[0]); j++ { // dp[0][j]
        if p[j-1] == '*' { dp[0][j] = dp[0][j-2] }
    }
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if isLowercase(p[j-1]) {
                dp[i][j] = s[i-1] == p[j-1] && dp[i-1][j-1]
            } else if p[j-1] == '*' { // j >= 2 if p is valid
                if s[i-1] == p[j-2] || p[j-2] == '.' {
                    dp[i][j] = dp[i-1][j] || dp[i][j-2]
                } else {
                    dp[i][j] = dp[i][j-2]
                }
            } else { // p[j-1] == '.'
                dp[i][j] = dp[i-1][j-1]
            }
        }
    }
    return dp[len(s)][len(p)]
}

func isLowercase(b byte) bool { return b >= 'a' && b <= 'z' }

// dp
// dp[len(s)+1][len(p)+1] []bool
// dp[i][j]: s[:i] p[:j] match?
// dp[0][0] = true
// dp[1..len(s)][0] = false
// dp[0][1..len(p)] uncertain. p: ".*" "a*"

// p must be valid "*"
// dp[i][j]: v = p[j-1] 
// if v a~z: s[i-1] == v && dp[i-1][j-1]
// if v '*': (j-2>=0) if s[i-1] == p[j-2] || p[j-2] == '.' { dp[i-1][j] || dp[i][j-2] }
//                    else { dp[i][j-2] }
// if v '.': dp[i-1][j-1]
