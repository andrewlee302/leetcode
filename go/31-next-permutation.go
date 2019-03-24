func nextPermutation(nums []int)  {
    if len(nums) == 0 || len(nums) == 1 { return }
    i := len(nums) - 1
    for ; i > 0 && nums[i] <= nums[i-1]; i-- { } // must be <= instead of <. E.g. [5,1,1].
    if i > 0 {
        j := i
        for ; j < len(nums); j++ {
            if nums[j] <= nums[i-1] { break } // must be <= instead of <. E.g. [1, 5, 1].
        }
        nums[j-1], nums[i-1] = nums[i-1], nums[j-1]
    }
    for left, right := i, len(nums) - 1; left < right; {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
