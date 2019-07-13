// DP.
func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1)+1)
    for i := 0; i < len(word1)+1; i++ { 
        dp[i] = make([]int, len(word2)+1) 
        for j := 0; j < len(word2)+1; j++ {
            if i == 0 || j == 0 { 
                dp[i][j] = i|j 
            } else {
                if word1[i-1] == word2[j-1] {
                    dp[i][j] = dp[i-1][j-1]
                } else {
                    dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
                }
            }
        }
    }
    return dp[len(word1)][len(word2)]
}

func min(i, j int) int { if i < j { return i } else { return j } }

// i: 0 ~ len(word1), j: 0 ~ len(word2)
// if w1[i] = w2[j]:
//  dp[i][j] = dp[i-1][j-1] + 1
// else:
//  dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
//      dp[i-1][j-1] + 1: replace w1[i] to w2[j]
//      dp[i-1][j] + 1: remove w1[i]
//      dp[i][j-1] + 1: add w2[j] to w1[:i+1]
// dp[i][0] = i
// dp[0][j] = j
// dp[0][0] = 0
