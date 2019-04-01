func findTargetSumWays(nums []int, S int) int {
    if len(nums) == 0 || S > 1000 || S < -1000 { return 0 }
    dp := make([][]int, len(nums))
    dp[0] = make([]int, 2001) 
    for j := 0; j < 2001; j++ { dp[0][j] = 0 }
    // dp[0][nums[0]+1000], dp[0][-nums[0]+1000] = 1, 1
    dp[0][nums[0]+1000]++
    dp[0][-nums[0]+1000]++
    for i := 1; i < len(nums); i++ {
        dp[i] = make([]int, 2001)
        for j := 0; j < 2001; j++ { dp[i][j] = -1 }
    }
    return helper(nums, S, dp, len(nums) - 1)
}

func helper(nums []int, target int, dp [][]int, i int) int {
    if dp[i][target+1000] != -1 { return dp[i][target+1000] }
    try1, try2 := target-nums[i], target+nums[i]
    ways := 0
    if try1 <= 1000 && try1 >= -1000 { ways += helper(nums, try1, dp, i-1) }
    if try2 <= 1000 && try2 >= -1000 { ways += helper(nums, try2, dp, i-1) }
    dp[i][target+1000] = ways
    return ways
}
