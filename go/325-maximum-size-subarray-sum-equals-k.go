func maxSubArrayLen(nums []int, k int) int {
    m := make(map[int]int)
    sum := 0
    m[0] = -1
    maxLen := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if idx, ok := m[sum-k]; ok {
            if i - idx > maxLen { maxLen = i - idx }
        }
        if _, ok := m[sum]; !ok { m[sum] = i }
    }
    return maxLen
}
