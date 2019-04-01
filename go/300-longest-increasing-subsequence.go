func lengthOfLIS(nums []int) int {
    if len(nums) == 0 { return len(nums) }
    dp := make([]int, len(nums))
    dp[0] = 1
    max := 1
    for i := 1; i < len(nums); i++ {
        iMax := 1 
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if iMax < dp[j] + 1 { iMax = dp[j] + 1 }
            }
        }
        dp[i] = iMax
        if iMax > max { max = iMax }
    }
    fmt.Println(dp)
    return max
}
