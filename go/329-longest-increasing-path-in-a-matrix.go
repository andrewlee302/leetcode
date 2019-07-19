// DFS + memorization
func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return 0 }
    moves := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    dp := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++ { dp[i] = make([]int, len(matrix[0])) }
    maxPath := 0
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            maxPath = max(maxPath, DFS(matrix, dp, moves, i, j))
        }
    }
    return maxPath
}

func DFS(matrix, dp, moves [][]int, i, j int) int {
    if dp[i][j] != 0 { return dp[i][j] }
    maxPath := 0
    for _, move := range moves {
        ni, nj := i+move[0], j+move[1]
        if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len(matrix[0]) && matrix[ni][nj] > matrix[i][j] {
            maxPath = max(maxPath, DFS(matrix, dp, moves, ni, nj))
        }
    }
    dp[i][j] = maxPath + 1
    return dp[i][j]
}

func max(i, j int) int { if i > j { return i } else { return j } }
