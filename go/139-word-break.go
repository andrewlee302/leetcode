func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]bool)
    for _, w := range wordDict { dict[w] = true }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            dp[i] = dp[j] && dict[s[j:i]]
            if dp[i] { break }
        }
    }
    return dp[len(s)]
}

// dp[0..len(s)] len(s) inclusive
// dp[0] = true
// dp[i] = dp[0] && dict[s[0:i]] || dp[1] && dict[s[1:i]] || xxx || dp[i-1] && dict[s[i-1:i]] 
