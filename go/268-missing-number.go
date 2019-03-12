// Gauss' Formula: O(N), O(1)
func missingNumber(nums []int) int {
    sum, expectedSum := 0, len(nums)*(len(nums)+1)/2
    for _, v := range nums {
        sum += v
    }
    return expectedSum - sum
}

// Bit Manipulation: O(N), O(1)
func missingNumber(nums []int) int {
    ret := len(nums)
    for i := 0; i < len(nums); i++ {
        ret ^= i ^ nums[i]
    }
    return ret
}
