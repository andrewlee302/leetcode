func islandPerimeter(grid [][]int) int {
    landCnt := 0
    touchCnt := 0 
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                landCnt++
                if j + 1 < len(grid[0]) && grid[i][j+1] == 1 { touchCnt++ }
                if i + 1 < len(grid) && grid[i+1][j] == 1 { touchCnt++ }
            }
        }
    }
    return landCnt*4 - touchCnt*2
}
