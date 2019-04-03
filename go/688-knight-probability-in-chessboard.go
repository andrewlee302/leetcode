func knightProbability(N int, K int, r int, c int) float64 {
    moves := [][]int{{1,2},{1,-2},{-1,2},{-1,-2},{2,1},{2,-1},{-2,1},{-2,-1}}
    dp := make([][][]float64, N)
    for i := 0; i < N; i++ { 
        dp[i] = make([][]float64, N) 
        for j := 0; j < N; j++ { 
            dp[i][j] = make([]float64, K+1) 
            dp[i][j][0] = 1.0 // Note
            for l := 1; l < K+1; l++ { dp[i][j][l] = -1 }
        }
    }
    forward(N, K, K, r, c, moves, dp) 
    return dp[r][c][K] // Note K == 0
}

func forward(N, cnt, K, r, c int, moves [][]int, dp [][][]float64) float64 {
    if cnt == 0 { return 1 }
    if dp[r][c][cnt] != -1 { return dp[r][c][cnt] }
    possibility := 0.0
    for _, move := range moves {
        rr, cc := r + move[0], c + move[1]
        if rr >= 0 && rr < N && cc >= 0 && cc < N { 
            possibility += forward(N, cnt-1, K, rr, cc, moves, dp) 
        }
    }
    possibility /= 8
    dp[r][c][cnt] = possibility
    return possibility
}
