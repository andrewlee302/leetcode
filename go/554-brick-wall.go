func leastBricks(wall [][]int) int {
    if len(wall) == 0 || len(wall[0]) == 0 { return 0 }
    m := make(map[int]int)
    for i := 0; i < len(wall); i++ {
        end := 0
        for j := 0; j < len(wall[i]) - 1; j++ {
            end += wall[i][j]
            m[end]++
        }
    }
    maxEndCnt := 0
    for _, cnt := range m { if cnt > maxEndCnt { maxEndCnt = cnt } }
    return len(wall) - maxEndCnt
}
