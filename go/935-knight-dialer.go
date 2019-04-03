const divisor = 1000000007
func knightDialer(N int) int {
    moveTos := [][]int{{4,6},{6,8},{7,9},{4,8},{3,9,0},{},{1,7,0},{2,6},{1,3},{2,4}}
    var dp [2][10]int
    for j := 0; j < 10; j++ { dp[0][j] = 1 } // initial
    for i := 0; i < N - 1; i++ {
        for j := 0; j < 10; j++ {
            sum := 0
            for _, moveTo := range moveTos[j] {
                sum += dp[i&1][moveTo]
                sum = sum%divisor
            }
            dp[(^i)&1][j] = sum
        }
    }
    sum := 0
    for j := 0; j < 10; j++ {
        sum += dp[(^N-2)&1][j]
        sum = sum%divisor
    }
    return sum
}
