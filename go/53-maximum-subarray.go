const int_max = int(^uint(0)>>1)
const int_min = ^int_max
func maxSubArray(nums []int) int {
    maxSum := int_min
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum > maxSum { maxSum = sum }
        if sum < 0 {
            sum = 0
        }
    }
    return maxSum
}


// dp version
func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    maxSum := dp[0]
    for i := 1; i < len(nums); i++ {
        if dp[i-1] < 0 { dp[i] = nums[i]  } else { dp[i] = dp[i-1] + nums[i] }
        if dp[i] > maxSum { maxSum = dp[i] }
    }
    return maxSum
}
