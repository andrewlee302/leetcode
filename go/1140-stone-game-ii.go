const INT_MAX = int(^uint(0)>>1)
func stoneGameII(piles []int) int {
    N := len(piles)
    dp := make([][][2]int, N+1)
    for i := 0; i < N+1; i++ { dp[i] = make([][2]int, N) }
    return helper(0, 1, 0, dp, piles)
}

func helper(i, m, turn int, dp [][][2]int, piles []int) int {
    if i == len(piles) || dp[i][m][turn] != 0 { return dp[i][m][turn] }
    var profit int
    if turn == 0 {
        pileSum := 0
        for x := 1; x <= min(2*m, len(piles)-i); x++ {
            pileSum += piles[i+x-1]
            profit = max(profit, pileSum + helper(i+x, max(m, x), 1-turn, dp, piles))
        }
    } else {
        profit = INT_MAX
        for x := 1; x <= min(2*m, len(piles)-i); x++ {
            profit = min(profit, helper(i+x, max(m, x), 1-turn, dp, piles))
        }
    }
    dp[i][m][turn] = profit
    return profit
}

func min(i, j int) int { if i < j { return i } else { return j } } 
func max(i, j int) int { if i > j { return i } else { return j } } 

/*
N = len(piles)
dp represents my profits in piles[i:].
dp[i][m][0] // my turn
dp[i][m][1] // other turn

recursive memorization is better than normal one.

dp[N][*][0], dp[N][*][0] = 0, 0
*/

