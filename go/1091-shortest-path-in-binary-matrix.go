// BFS level-order-traversal.
import "container/list"
func shortestPathBinaryMatrix(grid [][]int) int {
    N := len(grid)
    if grid[0][0] == 1 || grid[N-1][N-1] == 1 { return -1 }
    moves := [][]int{{0,1},{0,-1},{1,0},{-1,0},{1,1},{1,-1},{-1,1},{-1,-1}}
    queue := list.New()
    queue.PushBack([]int{0,0})
    pathLen := 1
    for queue.Len() != 0 {
        pathLen++
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            v := e.Value.([]int)
            for _, move := range moves {
                nX, nY := v[0] + move[0], v[1] + move[1]
                if nX == N-1 && nY == N-1 { return pathLen }
                if nX >= 0 && nX < N && nY >= 0 && nY < N && grid[nX][nY] == 0 {
                    grid[nX][nY] = 1
                    queue.PushBack([]int{nX, nY})
                }
            }
        }
    }
    return -1
}
