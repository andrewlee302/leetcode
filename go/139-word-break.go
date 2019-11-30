func wordBreak(s string, wordDict []string) bool {
    if len(s) == 0 { return false }
    dictM := make(map[string]bool)
    for _, str := range wordDict { dictM[str] = true }
    dp := make([]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = dictM[s[:i+1]]
        for j := i - 1; j >= 0; j-- {
            dp[i] = dp[i] || (dp[j] && dictM[s[j+1:i+1]])
        }
    }
    return dp[len(s)-1]
}
