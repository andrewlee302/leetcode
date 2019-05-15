func gardenNoAdj(N int, paths [][]int) []int {
    graph := make([][]int, N)
    for _, path := range paths {
        graph[path[0]-1] = append(graph[path[0]-1], path[1]-1)
        graph[path[1]-1] = append(graph[path[1]-1], path[0]-1)
    }
    ret := make([]int, N)
    for i := 0; i < N; i++ {
        nears := graph[i]
        var slots [5]bool
        for _, near := range nears {
            if near < i {
                slots[ret[near]] = true
            }
        }
        for j := 1; j < 5; j++ {
            if !slots[j] {
                ret[i] = j
                break
            }
        }
    }
    return ret
}
