func numDecodings(s string) int {
    mod := 1000000000 + 7
    if len(s) == 0 { return 0 }
    dp := make([]int, len(s)+1)
    dp[0] = 1
    if s[0] == '*' { dp[1] = 9 } else if s[0] == '0' { dp[1] = 0 } else { dp[1] = 1 } // Note!!!
    for i := 2; i < len(dp); i++ {
        si := i-1
        if s[si] == '*' {
            if s[si-1] == '1' {
                dp[i] = dp[i-1]*9 + dp[i-2]*9
            } else if s[si-1] == '2' {
                dp[i] = dp[i-1]*9 + dp[i-2]*6
            } else if s[si-1] == '*' { // Note!!!
                dp[i] = dp[i-1]*9 + dp[i-2]*15
            } else {
                dp[i] = dp[i-1]*9
            }
        } else if s[si] == '0' {
            if s[si-1] == '*' {
                dp[i] = dp[i-2]*2
            } else if s[si-1] >= '1' && s[si-1] < '3' { 
                dp[i] = dp[i-2] 
            } else { dp[i] = 0 }
        } else if s[si] >= '1' && s[si] <= '6' && s[si-1] == '*' { // Note!!!
            dp[i] = dp[i-1] + dp[i-2]*2
        } else if s[si] >= '1' && s[si] <= '6' && (s[si-1] == '1' || s[si-1] == '2') || s[si] >= '7' && s[si] <= '9' && (s[si-1] == '1' || s[si-1] == '*') {
            dp[i] = dp[i-1] + dp[i-2]
        } else {
            dp[i] = dp[i-1]
        }
        dp[i] = dp[i]%mod
    }
    return dp[len(s)]
}
// dp: 0~len(s) (inclusively)
// corner case
// dp[0] = 1
// dp[1]: if s[0] == '*': 9 else if s[0] = '0': 0 else: 1
//
// i = i-1
// if s[i] == 0
//      if s[i-1] == '*': dp[i] = dp[i-2]*2
//      else if s[i-1] < 3: dp[i] = dp[i-2]
//      else: dp[i] = 0
// else if s[i] == 1~6
//      if s[i-1] == '*': dp[i] = dp[i-1] + dp[i-2]*2
//      else if 0 < s[i-1] < 3: dp[i] = dp[i-1] + dp[i-2]
//      else: dp[i] = dp[i-1]
// else if s[i] == 7~9
//      if s[i-1] == '*': dp[i] = dp[i-1] + dp[i-2]
//      if 0 < s[i-1] < 2: dp[i] = dp[i-1] + dp[i-2]
//      else: dp[i] = dp[i-1]
// else if s[i] == '*'
//      if s[i-1] == 1: dp[i] = dp[i-1]*9 + dp[i-2]*9
//      else if s[i-1] == 2: dp[i] = dp[i-1]*9 + dp[i-2]*6
//      else if s[i-1] == *: dp[i] = dp[i-1]*9 + dp[i-2]*15
//      else: dp[i] = dp[i-1]*9

