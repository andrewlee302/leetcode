// DP, O(N^(3)), O(N^(2)).
const INT_MAX = int(^uint(0)>>1)
func mctFromLeafValues(arr []int) int {
    if len(arr) < 2 { return 0 }
    dp := make([][][2]int, len(arr))
    for i := 0; i < len(arr); i++ { 
        dp[i] = make([][2]int, len(arr))
        dp[i][i] = [2]int{arr[i], 0} 
    }
    for l := 2; l <= len(arr); l++ {
        for i := 0; i <= len(arr)-l; i++ {
            j := i+l-1
            dp[i][j][1] = INT_MAX
            for k := i; k < j; k++ {
                dp[i][j][1] = min(dp[i][j][1], dp[i][k][1]+dp[i][k][0]*dp[k+1][j][0]+dp[k+1][j][1])
            }
            dp[i][j][0] = max(dp[i][j-1][0], arr[j])
        }
    }
    return dp[0][len(arr)-1][1]
}

func max(i, j int) int { if i > j { return i } else { return j } }
func min(i, j int) int { if i < j { return i } else { return j } }

// dp[i][i] = []int{arr[i], 0}
// dp[i][j][1] = min(dp[i][k][1] + dp[i][k][0]*dp[k+1][j][0] +  dp[k+1][j][1]) k=i~(j-1)
// dp[i][j][0] = max(dp[i][j-1][0], arr[j])
