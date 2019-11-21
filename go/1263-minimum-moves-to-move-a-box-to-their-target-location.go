// Two-layer BFS, outer layer use two position as one element.
import "container/list"
func minPushBox(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 { return -1 }
    moves := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
    var pX, pY, bX, bY, tX, tY int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            switch (grid[i][j]) {
                case 'S': pX, pY = i, j
                case 'B': bX, bY = i, j
                case 'T': tX, tY = i, j
            }
        }
    }
    visited := make(map[string]bool)
    queue := list.New()
    queue.PushBack([2][2]int{{pX, pY}, {bX, bY}})
    cnt := 0
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            pos := e.Value.([2][2]int)
            if pos[1][0] == tX && pos[1][1] == tY { return cnt }
            for _, move := range moves {
                pXX, pYY := pos[1][0] - move[0], pos[1][1] - move[1]
                bXX, bYY := pos[1][0] + move[0], pos[1][1] + move[1]
                s := toString(pXX, pYY, bXX, bYY)
                if bXX >= 0 && bXX < len(grid) && bYY >= 0 && bYY < len(grid[0]) && grid[bXX][bYY] != '#' && !visited[s] && isAccessible(grid, moves, pos[0][0], pos[0][1], pXX, pYY, pos[1][0], pos[1][1]) {
                    visited[s] = true
                    queue.PushBack([2][2]int{{pXX, pYY}, {bXX, bYY}})
                }
            }
        }
        cnt++
    }
    return -1
}

func toString(x1, y1, x2, y2 int) string {
    return strconv.Itoa(x1) + "," + strconv.Itoa(y1) + "," + strconv.Itoa(x2) + "," + strconv.Itoa(y2)
}

func isAccessible(grid [][]byte, moves [][]int, pX, pY, pXX, pYY, bX, bY int) bool {
    if pX < 0 || pX >= len(grid) || pY < 0 || pY >= len(grid[0]) || grid[pX][pY] == '#' { return false }
    if pXX < 0 || pXX >= len(grid) || pYY < 0 || pYY >= len(grid[0]) || grid[pXX][pYY] == '#' { return false }
    visited := make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ { visited[i] = make([]bool, len(grid[0])) }
    queue := list.New()
    queue.PushBack([]int{pX, pY})
    visited[pX][pY] = true
    for queue.Len() != 0 {
        e := queue.Front()
        queue.Remove(e)
        pos := e.Value.([]int)
        if pos[0] == pXX && pos[1] == pYY { return true }
        for _, move := range moves {
            nextX, nextY := pos[0] + move[0], pos[1] + move[1]
            if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] != '#' && (nextX != bX || nextY != bY) && !visited[nextX][nextY] {
                visited[nextX][nextY] = true
                queue.PushBack([]int{nextX, nextY})
            }
        }
    }
    return false
} 
