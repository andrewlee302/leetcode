import "container/list"
func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
    matrix := make([][]bool, 2*n) // r: 0-(n-1), b: n-(n*2-1)
    for i := 0; i < 2*n; i++ { matrix[i] = make([]bool, 2*n) }
    for _, rEdge := range red_edges { matrix[rEdge[0]][rEdge[1]+n] = true }
    for _, bEdge := range blue_edges { matrix[bEdge[0]+n][bEdge[1]] = true }
    visited := make([]bool, 2*n)
    path := make([]int, n)
    for i := 0; i < n; i++ { path[i] = -1 }
    queue := list.New()
    queue.PushBack(0)
    queue.PushBack(n)
    visited[0], visited[n] = true, true
    level := 0
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            v := e.Value.(int)
            if v < n && path[v] == -1 {
                path[v] = level
            } else if v >= n && path[v-n] == -1 {
                path[v-n] = level
            }
            queue.Remove(e)
            for j := 0; j < 2*n; j++ {
                if matrix[v][j] && !visited[j] {
                    queue.PushBack(j)
                    visited[j] = true
                }
            }
        }
        level++
    }
    return path
}
