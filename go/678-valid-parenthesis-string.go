// greedy
func checkValidString(s string) bool {
    lower, upper := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' { lower++ } else { lower-- }
        if s[i] != ')' { upper++ } else { upper-- }
        if upper < 0 { return false }
        if lower < 0 { lower = 0 } // can't below zero.
    }
    return lower == 0
}

// dp
func checkValidString(s string) bool {
    if len(s) == 0 { return true }
    dp := make([][]bool, len(s)+1)
    for i := 0; i < len(dp); i++ { dp[i] = make([]bool, len(s)+1) }
    dp[0][0] = true
    for i := 1; i < len(dp); i++ {
        var moves []int
        if s[i-1] == '(' { 
            moves = []int{-1}
        } else if s[i-1] == ')' { 
            moves = []int{1}
        } else {
            moves = []int{1, 0, -1}
        }
        for j := 0; j < len(s) + 1; j++ {
            flag := false
            for _, move := range moves {
                jj := j + move 
                if jj >= 0 && jj < len(s)+1 {
                    flag = flag || dp[i-1][jj]  
                }       
            }
            dp[i][j] = flag   
        }
    }
    return dp[len(s)][0]
}
