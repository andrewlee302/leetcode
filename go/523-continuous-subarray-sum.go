// Store location.
func checkSubarraySum(nums []int, k int) bool {
    if k < 0 { k = -k } 
    m := make(map[int]int) // map[sum](1st pos)
    sum := 0
    m[0] = -1
    for i := 0; i < len(nums); i++ {
        var remainder int
        sum += nums[i]
        if k == 0 { remainder = sum } else { remainder = sum%k }
        if pos, ok := m[remainder]; ok && i - pos >= 2 {
            return true
        } else if !ok {
						m[remainder] = i
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
