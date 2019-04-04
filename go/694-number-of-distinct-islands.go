import "strings"
func numDistinctIslands(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 { return 0 }
    moves := [][]int{{1,0},{0,1},{0,-1},{-1,0}}
    actions := []byte{'d','r','l','u'} // 3-i can get the opposite direction.
    islands := make(map[string]bool)
    var sb strings.Builder
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                grid[i][j] = -1
                DFS(i, j, grid, moves, actions, &sb)
                island := sb.String()
                sb.Reset()
                fmt.Println(island)
                islands[island] = true
            }
        }
    }
    return len(islands)
}

func DFS(r, c int, grid, moves [][]int, actions []byte, sb *strings.Builder) {
    for i, move := range moves {
        nextR, nextC := r + move[0], c + move[1]
        if nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0]) && grid[nextR][nextC] == 1 {
            grid[nextR][nextC] = -1
            sb.WriteByte(actions[i])
            DFS(nextR, nextC, grid, moves, actions, sb)
            sb.WriteByte(3 - actions[i]) // Note: retreat path can distinguish the corner case.
        }
    }
}
