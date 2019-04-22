import "container/list"
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    var ret [][]int
    moves := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    visited := make([][]bool, R)
    for i := 0; i < R; i++ { visited[i] = make([]bool, C) }
    queue := list.New()
    queue.PushBack([]int{r0, c0})
    visited[r0][c0] = true
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            pos := e.Value.([]int)
            ret = append(ret, pos)
            for _, move := range moves {
                r1, c1 := pos[0] + move[0], pos[1] + move[1]
                if r1 >= 0 && r1 < R && c1 >= 0 && c1 < C && !visited[r1][c1] {
                    visited[r1][c1] = true
                    queue.PushBack([]int{r1,c1})
                }
            }
        }
    }
    return ret
}
