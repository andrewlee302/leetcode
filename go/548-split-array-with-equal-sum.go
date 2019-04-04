func splitArray(nums []int) bool {
    if len(nums) < 7 { return false }
    sums := make([]int, len(nums))
    for sum, i := 0, 0; i < len(nums); i++ { 
        sum += nums[i]
        sums[i] = sum
    }
    for j := 3; j < len(nums) - 3; j++ {
        tryI := make(map[int]bool)
        for i := 1; i <= j - 2; i++ {
            if sums[i-1] == sums[j-1] - sums[i] { tryI[sums[i-1]] = true }
        }
        for k := j + 2; k < len(nums) - 1; k++ {
            sum_j_k := sums[k-1] - sums[j]
            sum_k := sums[len(nums)-1] - sums[k]
            if sum_j_k == sum_k && tryI[sum_k] { return true }
        }
    }
    return false
}
