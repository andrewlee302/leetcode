// Forward scan.
func maxProfit(prices []int) int {
    if len(prices) == 0 { return 0 }
    ret := 0
    minPrice := prices[0]
    for _, price := range prices[1:] {
        ret = max(ret, price-minPrice)
        minPrice = min(minPrice, price)
    }
    return ret
}

// Reverse scan.
func maxProfit(prices []int) int {
    if len(prices) == 0 { return 0 }
    ret := 0
    maxPrice := prices[len(prices)-1]
    for i := len(prices) - 2; i >= 0; i-- {
        ret = max(ret, maxPrice - prices[i])
        maxPrice = max(maxPrice, prices[i])
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }
func min(i, j int) int { if i < j { return i } else { return j } }
