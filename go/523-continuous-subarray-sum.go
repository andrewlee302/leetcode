// Store location.
func checkSubarraySum(nums []int, k int) bool {
    if k < 0 { k = -k } // Note.
    m := make(map[int]int) // map[sum](1st pos)
    remainder := 0
    m[0] = -1 // Note.
    for i := 0; i < len(nums); i++ {
        remainder += nums[i]
        if k != 0 { remainder = remainder%k }
        if pos, ok := m[remainder]; !ok {
            m[remainder] = i
        } else if ok && i - pos >= 2 {
            return true
        }
    }
    return false
}

// No need to store locations.
func checkSubarraySum(nums []int, k int) bool {
    if k < 0 { k = -k }
    remainder := 0
    m := make(map[int]bool)
    prevRemainder := 0
    for _, num := range nums {
        remainder += num
        if k != 0 { remainder = remainder%k }
        if m[remainder] {
            return true
        } else {
            m[prevRemainder] = true
            prevRemainder = remainder
        }
    }
    return false
}
