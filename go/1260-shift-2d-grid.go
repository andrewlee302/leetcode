func shiftGrid(grid [][]int, k int) [][]int {
    if len(grid) == 0 || len(grid[0]) == 0 { return nil }
    ret := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ { ret[i] = make([]int, len(grid[0]))}
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            absIdx := i * len(grid[0]) + j
            nextAbsIdx := (absIdx + k) % (len(grid) * len(grid[0]))
            nextI, nextJ := nextAbsIdx/len(grid[0]), nextAbsIdx%len(grid[0])
            ret[nextI][nextJ] = grid[i][j]
        }
    }
    return ret
}
