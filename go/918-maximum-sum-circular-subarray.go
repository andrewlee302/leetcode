const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func maxSubarraySumCircular(A []int) int {
    if len(A) == 0 { return 0 }
    min, max := INT_MAX, INT_MIN
    minTry, maxTry, sum := 0, 0, 0
    for i := 0; i < len(A); i++ {
        sum += A[i]
        minTry, maxTry = minTry + A[i], maxTry + A[i]
        if maxTry > max { max = maxTry }
        if maxTry < 0 {  maxTry = 0 }
        if minTry < min { min = minTry }
        if minTry > 0 { minTry = 0 }
    }
    // Note: [-2,-3,-1]
    if max < 0 { return max }
    if max > sum - min { return max } else { return sum - min }
}
