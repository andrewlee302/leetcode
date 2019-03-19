func findLengthOfLCIS(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    start := 0
    maxCnt := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i-1] {
            if i - start > maxCnt {
                maxCnt = i - start
            }
            start = i
        }
    }
    if len(nums) - start > maxCnt {
        maxCnt = len(nums) - start
    }
    return maxCnt
}
