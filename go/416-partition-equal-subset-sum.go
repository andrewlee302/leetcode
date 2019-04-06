// normal bottom-up dp.
func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums { sum += num }
    if sum%2 != 0 { return false }
    target := sum/2
    dp := make([][]bool, len(nums)+1)
    for i := 0; i < len(nums)+1; i++ { dp[i] = make([]bool, target+1) }
    dp[0][0] = true
    for i := 1; i < len(nums)+1; i++ {
        for j := 0; j < target+1; j++ {
            dp[i][j] = dp[i-1][j]
            if j - nums[i-1] >= 0 {
                dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
            }
        }
        if dp[i][target] == true { return true }
    }
    return false
}

// optimization for space.
func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums { sum += num }
    if sum%2 != 0 { return false }
    target := sum/2
    dp := make([]bool, target+1)
    dp[0] = true
    for i := 1; i < len(nums)+1; i++ {
        for j := target; j >= 0; j-- { // must visit from highest to lowest.
            if j - nums[i-1] >= 0 {
                dp[j] = dp[j] || dp[j-nums[i-1]]
            }
        }
        if dp[target] == true { return true }
    }
    return false
}

// dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
// dp[j] = dp[j-nums[i]] || dp[j] // must from the highest to lowest
