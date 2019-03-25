import "container/list"

const (
    EMPTY = 1 << 31 - 1
    WALL = -1
    GATE = 0
)

func wallsAndGates(rooms [][]int)  {
    if len(rooms) == 0 || len(rooms[0]) == 0 { return }
    moves := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
    m, n := len(rooms), len(rooms[0])
    queue := list.New()
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == GATE { queue.PushBack([]int{i, j})}
        }
    }
    for queue.Len() != 0 {
        e := queue.Front()
        queue.Remove(e)
        pos := e.Value.([]int)
        for _, move := range moves {
            x := pos[0] + move[0]
            y := pos[1] + move[1]
            if x >= m || y >= n || x < 0 || y < 0 || rooms[x][y] != EMPTY { continue }
            rooms[x][y] = rooms[pos[0]][pos[1]] + 1
            queue.PushBack([]int{x, y})
        }
    }
}
