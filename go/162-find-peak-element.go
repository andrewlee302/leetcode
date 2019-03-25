func findPeakElement(nums []int) int {
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i+1] { return i }
    }
    return len(nums) - 1
}
