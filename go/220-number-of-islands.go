func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    cnt := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                cnt++
                DFS(grid, i, j)
            }
        }
    }
    return cnt
}

// i,j are valid and grid[i][j] muse be 1
func DFS(grid [][]byte, i, j int) {
    grid[i][j] = '0'
    coords := []int{i+1,j,i-1,j,i,j+1,i,j-1}
    for k := 0; k <len(coords) ; k+=2 {
        x, y := coords[k], coords[k+1]
        if x >=0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1'{
            DFS(grid, x, y)
        }
    }
}
