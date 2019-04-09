// linear scan O(N)
func findPeakElement(nums []int) int {
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i+1] { return i }
    }
    return len(nums) - 1
}

// quickselect O(logN)
func findPeakElement(nums []int) int {
    return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, l, r int) int {
    if l == r { return l }
    mid := l + (r-l)/2
    if nums[mid] > nums[mid+1] {
        return helper(nums, l, mid)
    } else {
        return helper(nums, mid+1, r)
    }
}
