// greedy
const INT_MAX = int(^uint(0)>>1)
const INT_MIN = ^INT_MAX
func maxProfit(prices []int) int {
    if len(prices) == 0 { return 0 }
    firstBuy, firstSell, secondBuy, secondSell := INT_MIN, 0, INT_MIN, 0
    for _, price := range prices {
        secondSell = max(secondSell, price+secondBuy)
        secondBuy = max(secondBuy, firstSell-price)
        firstSell = max(firstSell, price+firstBuy)
        firstBuy = max(firstBuy, -price)
    }
    return max(firstSell, secondSell) // secondSell mustn't be less than firstSell
}

func max(a, b int) int { if a > b { return a } else { return b } }

// dp
func maxProfit(prices []int) int {
    if len(prices) == 0 { return 0 }
    dp := make([][]int, 2)
    txnLimit := 2
    for i := 0; i < 2; i++ { dp[i] = make([]int, len(prices)) }
    for k := 1; k <= txnLimit; k++ {
        kk := k&1
        maxRest := -prices[0]
        for i := 1; i < len(prices); i++ {
            dp[kk][i] = max(dp[kk][i-1], prices[i] + maxRest)
            maxRest = max(maxRest, dp[1-kk][i-1] - prices[i])
        }
    }
    return dp[txnLimit&1][len(prices)-1]
}

func max(a, b int) int { if a > b { return a } else { return b } }
// We can assume: sell and buy the same day is forbidden.
// dp[k][i]: max profit at most k txns for 0~ith days.
// dp[k][i] = max(dp[k][i-1], max(prices[i] - prices[ii] + dp[k-1][ii-1])) ii: 0~(i-1)
//
// dp[0][X] = 0
// dp[k][0] = 0
// dp[1][X] = max(dp[1][X-1], max(prices[X] - prices[ii])) ii:0~(X-1)
// dp[2][X] = max(dp[2][X-1], max(prices[X] - prices[ii] + dp[1][ii-1])) ii:0~(X-1)
// optimization:
// 1. because dp[k][i] just has relations with dp[k-1][ii] instead of arbitrary dp,
// so just two array is enough.
// 2. max(prices[i] - prices[ii] + dp[k-1][ii-1]) could calculated with increasing i.
