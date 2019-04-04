// greedy


// dp
const INT_MAX = int(^uint(0)>>1)
func maxProfit(prices []int) int {
    if len(prices) == 0 { return 0 }
    // sell on j, buy on j. I.e., not a trade. Just skip that day.
    // dp[k, i] = max(dp[k, i-1], prices[i] - min(prices[j] - dp[k-1, j])) 
    trans := 2
    dp := make([][]int, 2) // using the two as dp[k] and dp[k-1]
    for i := 0; i < len(dp); i++ { dp[i] = make([]int, len(prices)) }
    for k := 1; k <= trans; k++ {
        minCost := INT_MAX
        for i := 1; i < len(prices); i++ {
            cost := prices[i-1] - dp[(^k)&1][i-1]
            if cost < minCost { minCost = cost }
            dp[k&1][i] = prices[i] - minCost
            if dp[k&1][i-1] > dp[k&1][i] { dp[k&1][i] = dp[k&1][i-1] }
        }
    }
    return dp[trans&1][len(prices)-1]
}
