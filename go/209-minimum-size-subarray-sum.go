func minSubArrayLen(s int, nums []int) int {
    left, right := 0, 0
    sum := 0
    minLen := len(nums) + 1
    for right < len(nums){
        sum += nums[right]
        for sum >= s { 
            // left mustn't be more than right
            sum -= nums[left] 
            if sum < s {
                l := right - left + 1
                if l < minLen { minLen = l }
            }
            left++
        }
        right++
    }
    if minLen == len(nums) + 1 { return 0 } else { return minLen }
}
