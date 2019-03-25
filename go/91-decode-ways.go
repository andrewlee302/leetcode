// resursive
func numDecodings(s string) int {
    if len(s) == 0 || len(s) == 1 && s[0] != '0' { return 1 }
    if s[0] == '0' { return 0 }
    if s[0] > '2' || s[0] == '2' && s[1] > '6' { 
        if s[1] == '0' { return 0 } else { return numDecodings(s[1:]) }
    }
    // below is 10-26 (inclusive)
    if s[1] == '0' { return numDecodings(s[2:]) } // 10, 20
    return numDecodings(s[1:]) + numDecodings(s[2:])
}

// iterative
func numDecodings(s string) int {
    if len(s) == 0 { return 1 }
    if s[0] == '0' { return 0 }
    dp := make([]int, len(s) + 1)
    dp[len(s)] = 1
    if s[len(s)-1] == '0' { dp[len(s)-1] = 0 } else { dp[len(s)-1] = 1 }
    for i := len(s) - 2; i >= 0; i-- {
        if s[i] == '0' { 
            dp[i] = 0
            continue
        }
        if s[i] > '2' || s[i] == '2' && s[i+1] > '6' { 
            if s[i+1] == '0' { dp[i] = 0 } else { dp[i] = dp[i+1] }
            continue
        }
        // below is 10-26 (inclusive)
        if s[i+1] == '0' { 
            dp[i] = dp[i+2]  // 10, 20
        } else {
            dp[i] = dp[i+1] + dp[i+2]
        } 
    }
    return dp[0]   
}
