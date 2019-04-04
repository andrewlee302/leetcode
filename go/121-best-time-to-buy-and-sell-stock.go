func maxProfit(prices []int) int {
    minPrice := int((^uint(0))>>1)
    maxProfit := 0
    for _, v := range prices {
        if v - minPrice > maxProfit {
            maxProfit = v - minPrice
        }
        if v < minPrice {
            minPrice = v
        }
    }
    return maxProfit
}
