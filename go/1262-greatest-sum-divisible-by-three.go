func maxSumDivThree(nums []int) int {
    dp := []int{0, 0, 0} // only dp[0] is valid as dp[i]%3 == i.
    for _, num := range nums {
        var temp [3]int
        for i := 0; i < 3; i++ {
            j := (i+3-num%3)%3
            if dp[j]%3 == j { // Note: it's important.
                temp[i] = max(dp[i], dp[j]+num)
            } else {
                temp[i] = dp[i]
            }
        }
        dp = temp[:]
    }
    return dp[0]
}

func max(i, j int) int { if i > j { return i } else { return j } }
