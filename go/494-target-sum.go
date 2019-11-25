// DP iteration.
func findTargetSumWays(nums []int, S int) int {
    if len(nums) == 0 || S > 1000 || S < -1000 { return 0 }
    dp := make([][]int, len(nums)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, 2001)
    }
    dp[0][1000] = 1
    for i := 1; i < len(nums) + 1; i++ {
        for j := 0; j < 2001; j++ { // v + 1000
            v1, v2 := j + nums[i-1], j - nums[i-1]
            if v1 >= 0 && v1 <= 2000 { dp[i][j] += dp[i-1][v1] }
            if v2 >= 0 && v2 <= 2000 { dp[i][j] += dp[i-1][v2] }
        }
    }
    return dp[len(nums)][S+1000]
}

// DP recursion.
func findTargetSumWays(nums []int, S int) int {
    if len(nums) == 0 || S > 1000 || S < -1000 { return 0 }
    dp := make([][]int, len(nums)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, 2001)
        if i > 0 { for j := 0; j < 2001; j++ { dp[i][j] = -1 } } // Note!!!
    }
    dp[0][1000] = 1
    return recursion(nums, S+1000, dp, len(nums) - 1)
}

func recursion(nums []int, target int, dp [][]int, i int) int {
    if dp[i+1][target] != -1 { return dp[i+1][target] }
    v1, v2 := target + nums[i], target - nums[i]
    ret := 0
    if v1 >= 0 && v1 <= 2000 { ret += recursion(nums, v1, dp, i-1) }
    if v2 >= 0 && v2 <= 2000 { ret += recursion(nums, v2, dp, i-1)  }
    dp[i+1][target] = ret
    return ret
}
