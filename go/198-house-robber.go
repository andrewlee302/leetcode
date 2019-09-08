func rob(nums []int) int {
    if len(nums) == 0 { return 0 } else if len(nums) == 1 { return nums[0] }
    x, y := 0, nums[0]
    for _, num := range nums[1:] {
        x, y = y, max(num+x, y)
    }
    return y
}

func max(i, j int) int { if i > j { return i } else { return j } }
// i: 0~len(nums)
// dp[i] = max(nums[i-1] + dp[i-2], dp[i-1])
// dp[0] = 0
// dp[1] = 1
