// prefix sum + DP + (split 3 parts)
func maxSumOfThreeSubarrays(nums []int, k int) []int {
    // max sum and the start pos of the subarray before i (inclusively).
    leftDP := make([][]int, len(nums)) 
    // max sum and the start pos of the subarray after i (inclusively).
    rightDP := make([][]int, len(nums)) 
    
    prefixSum := make([]int, len(nums)+1) // i -> i+1.
    for i := 0; i < len(nums); i++ {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }
    
    maxLeft := []int{0, -1} // sum & start pos of the first array
    for i := k - 1; i < len(nums); i++ {
        if prefixSum[i+1] - prefixSum[i-k+1] > maxLeft[0] {
            maxLeft = []int{prefixSum[i+1] - prefixSum[i-k+1], i-k+1}
        }
        leftDP[i] = maxLeft
    } 
    
    maxRight := []int{0, -1} // sum & start pos of the third array
    for i := len(nums) - k; i >= 0; i-- {
        // using >= because we want the lexicographically smallest one. 
        if prefixSum[i+k] - prefixSum[i] >= maxRight[0] {
            maxRight = []int{prefixSum[i+k] - prefixSum[i], i}
        }
        rightDP[i] = maxRight
    }
    
    maxTotal := []int{0, -1} // sum & start pos of the second array
    for i := k; i <= len(nums) - 2*k; i++ {
        leftSum := leftDP[i-1][0]
        rightSum := rightDP[i+k][0]
        midSum := prefixSum[i+k] - prefixSum[i]
        total := leftSum + rightSum + midSum
        if total > maxTotal[0] {
            maxTotal = []int{total, i}
        }
    }
    midIdx := maxTotal[1]
    return []int{leftDP[midIdx-1][1],maxTotal[1], rightDP[midIdx+k][1]}
}
