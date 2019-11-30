// DP.
func longestArithSeqLength(A []int) int {
    if len(A) == 0 { return 0 }
    dp := make([]map[int]int, len(A))
    for i := 0; i < len(A); i++ { dp[i] = make(map[int]int) }
    dp[0][0] = 1
    ret := 1
    for i := 1; i < len(A); i++ {
        dp[i][0] = 1
        for j := 0; j < i; j++ {
            diff := A[i] - A[j]
            v := max(dp[j][diff], 1)
            dp[i][diff] = max(dp[i][diff], v+1)
            ret = max(ret, dp[i][diff])
        }
    }
    return ret
}

func max(a, b int) int { if a > b { return a } else { return b } }

// brute-force.
func longestArithSeqLength(A []int) int {
    longestLen := 0
    for i := 0; i < len(A); i++ {
        for j := i+1; j < len(A); j++ {
            // A[i], A[j] ...
            diff := A[j] - A[i]
            cnt := 2
            prev := A[j]
            for k := j+1; k < len(A); k++ {
                if diff == A[k] - prev {
                    cnt++
                    prev = A[k]
                }
            }
            if cnt > longestLen { longestLen = cnt }
        }
    }
    return longestLen
}
