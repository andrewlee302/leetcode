func canPartition(nums []int) bool {
    dp := make([][]int, len(nums)) // 0: unset, 1 true, 2 false
    sum := 0
    for _, num := range nums { sum += num }
    if sum%2 != 0 { return false }
    target := sum/2
    for i := 0; i < len(nums); i++ { dp[i] = make([]int, target+1) }
    return helper(nums, dp, target, 0)
}

func helper(nums []int, dp [][]int, target, idx int) bool {
    if target == 0 { return true }
    if idx == len(nums) { if target == 0 { return true } else { return false } }
    if dp[idx][target] != 0 { return dp[idx][target] == 1 }
    ret := helper(nums, dp, target, idx+1)
    if target >= nums[idx] {
        ret = ret || helper(nums, dp, target-nums[idx], idx+1)    
    }
    if ret == true { dp[idx][target] = 1 } else { dp[idx][target] = 2 }
    return ret
}
