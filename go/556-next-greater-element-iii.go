func nextGreaterElement(n int) int {
    var nums []int
    for n > 0 {
        remainder := n%10
        nums = append(nums, remainder)
        n /= 10
    }
    if len(nums) == 1 { return -1 } // Single significant bit.
    // reverse 
    for l, r := 0, len(nums) - 1; l < r; l, r = l+1, r-1 { nums[l], nums[r] = nums[r], nums[l] }
    i := len(nums) - 1
    for ; i >= 1 && nums[i] <= nums[i-1]; i-- { } // choose the nums[i] bigger than nums[i-1]
    if i == 0 { return -1 } // 21
    j := len(nums) - 1
    for ; j >= i && nums[j] <= nums[i-1]; j-- { } // choose the nums[j] bigger than nums[i-1]
    nums[i-1], nums[j] = nums[j], nums[i-1]
    // reverse
    for l, r := i, len(nums) - 1; l < r; l, r = l+1, r-1 { nums[l], nums[r] = nums[r], nums[l] }
    const INT_MAX int32 = int32(^uint32(0)>>1)
    sum := 0
    for k := 0; k < len(nums); k++ { 
        if int32(sum) < (INT_MAX - int32(nums[k])) / 10 {
            sum = sum * 10 + nums[k]
        } else {
            return -1
        }
    }
    return sum
}
