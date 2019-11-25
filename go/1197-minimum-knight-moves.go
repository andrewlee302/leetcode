import "container/list"
import "strconv"
func minKnightMoves(x int, y int) int {
    if x == 0 && y == 0 { return 0 }
    if x < 0 { x = -x }
    if y < 0 { y = -y }
    moves := [][]int{{1,2},{2,1},{-1,-2},{-2,-1},{1,-2},{-1,2},{-2,1},{2,-1}}
    visited := make(map[string]bool)
    queue := list.New()
    queue.PushBack([]int{0,0})
    visited["0,0"] = true
    level := 0
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            pos := e.Value.([]int)
            for _, move := range moves {
                nextX, nextY := pos[0] + move[0], pos[1] + move[1]
                str := strconv.Itoa(nextX) + "," + strconv.Itoa(nextY)
                if nextX >= -1 && nextY >= -1 && !visited[str] { // NOTE!!! nextX >= -1 && nextY >= -1
                    if nextX == x && nextY == y { return level + 1 }
                    visited[str] = true
                    queue.PushBack([]int{nextX,nextY})
                }
            }
        }
        level++
    }
    return -1 // impossible.
}
