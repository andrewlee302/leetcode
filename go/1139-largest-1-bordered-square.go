func largest1BorderedSquare(grid [][]int) int {
    gridD, gridR := make([][]int, len(grid)), make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        gridD[i], gridR[i] = make([]int, len(grid[0])), make([]int, len(grid[0]))
    }
    for i := 0; i < len(grid); i++ {
        for j := len(grid[i])-1; j >= 0; j-- {
            if grid[i][j] == 1 {
                if j == len(grid[i]) - 1 {
                    gridR[i][j] = 1
                } else {
                    gridR[i][j] = gridR[i][j+1] + 1
                }
            }
        }
    }
    for j := 0; j < len(grid[0]); j++ {
        for i := len(grid)-1; i >= 0; i-- {
            if grid[i][j] == 1 {
                if i == len(grid) - 1 {
                    gridD[i][j] = 1
                } else {
                    gridD[i][j] = gridD[i+1][j] + 1
                }
            }
        }
    }
    maxSideLen := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            possSide := min(gridR[i][j], gridD[i][j])
            if possSide > maxSideLen {
                for k := maxSideLen+1; k <= possSide; k++ {
                    if gridR[i+k-1][j] >= k && gridD[i][j+k-1] >= k {
                        maxSideLen = k
                    }
                }
            }
        }
    }
    return maxSideLen * maxSideLen
}

func min(i, j int) int { if i < j { return i } else { return j } }
