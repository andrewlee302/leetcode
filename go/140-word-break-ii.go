// Iterative DP: possbile MLE, TLE.
import "strings"
func wordBreak(s string, wordDict []string) []string {
    if len(s) == 0 { return nil }
    dictM := make(map[string]bool)
    for _, str := range wordDict { dictM[str] = true }
    dp := make([][]string, len(s))
    var builder strings.Builder
    for i := 0; i < len(s); i++ {
        if dictM[s[:i+1]] { dp[i] = append(dp[i], s[:i+1]) }
        for j := i - 1; j >= 0; j-- {
            if dictM[s[j+1:i+1]] {
                for _, str := range dp[j] {
                    builder.WriteString(str)
                    builder.WriteString(" ")
                    builder.WriteString(s[j+1:i+1])
                    dp[i] = append(dp[i], builder.String())
                    builder.Reset()
                }
            }
        }
    }
    return dp[len(s)-1]
}

// Recursive DP.
import "strings"
func wordBreak(s string, wordDict []string) []string {
    if len(s) == 0 { return nil }
    dictM := make(map[string]bool)
    for _, str := range wordDict { dictM[str] = true }
    dp := make(map[int][]string)
    return recursion(s, len(s)-1, dictM, dp)
}

func recursion(s string, idx int, dictM map[string]bool, dp map[int][]string) []string {
    if idx == -1 { return nil }
    if v, ok := dp[idx]; ok { return v }
    var ret []string
    var builder strings.Builder
    for i := idx; i >= 0; i-- {
        if dictM[s[i:idx+1]] {
            if i == 0 {
                ret = append(ret, s[i:idx+1])
                continue
            }
            for _, str := range recursion(s, i-1, dictM, dp) {
                builder.WriteString(str)
                builder.WriteString(" ")
                builder.WriteString(s[i:idx+1])
                ret = append(ret, builder.String())
                builder.Reset()
            }
        }
    }
    dp[idx] = ret
    return ret
}

