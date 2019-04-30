import "container/list"
func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
    var toColor [][]int
    visited := make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ { visited[i] = make([]bool, len(grid[0])) }
    moves := [][]int{{0,1},{1,0},{-1,0},{0,-1}}
    queue := list.New()
    visited[r0][c0] = true
    originColor := grid[r0][c0]
    queue.PushBack([]int{r0,c0})
    for queue.Len() != 0 {
        e := queue.Front()
        queue.Remove(e)
        pos := e.Value.([]int)
        for _, move := range moves {
            r1, c1 := pos[0] + move[0], pos[1] + move[1]
            // boundary
            if !(r1 >= 0 && r1 < len(grid) && c1 >= 0 && c1 < len(grid[0])) || grid[r1][c1] != originColor {
                toColor = append(toColor, []int{pos[0], pos[1]})
            } else if !visited[r1][c1] {
                visited[r1][c1] = true
                queue.PushBack([]int{r1,c1})
            }
        }
    }
    for _, pos := range toColor { grid[pos[0]][pos[1]] = color }
    return grid
}
