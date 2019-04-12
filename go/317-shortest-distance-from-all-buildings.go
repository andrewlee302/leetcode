import "container/list"
func shortestDistance(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 { return -1 }
    moves := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    total := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        total[i] = make([]int, len(grid[i]))
        copy(total[i], grid[i])
    }
    tag := 0
    for i := 0; i < len(grid); i++ { 
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                level := 0
                queue := list.New()
                queue.PushBack([]int{i,j})
                for queue.Len() != 0 {
                    size := queue.Len()
                    for k := 0; k < size; k++ {
                        e := queue.Front()
                        queue.Remove(e)
                        pos := e.Value.([]int)
                        for _, move := range moves {
                            x, y := pos[0]+move[0], pos[1]+move[1]
                            if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == tag {
                                grid[x][y] = tag - 1
                                total[x][y] += level + 1
                                queue.PushBack([]int{x,y})
                            }
                        }
                    }
                    level++
                }
                tag--
            }
        }
    }
    minTotalDist := -1
    for i := 0; i < len(grid); i++ { 
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == tag {
                if minTotalDist == -1 || total[i][j] < minTotalDist {
                    minTotalDist = total[i][j]
                }
            }
        }
    }
    return minTotalDist
}
