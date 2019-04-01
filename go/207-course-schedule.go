func canFinish(numCourses int, prerequisites [][]int) bool {
    if numCourses == 0 || len(prerequisites) == 0 { return true }
    edges := make([][]int, numCourses)
    for _, pair := range prerequisites {
        edges[pair[1]] = append(edges[pair[1]], pair[0])
    }
    visited := make([]int, numCourses) // 0: not visited, 1: visit, but not finish 2: visited
    for i := 0; i < len(visited); i++ {
        if visited[i] == 0 { 
            if !DFS(i, visited, edges) { return false }
        }
    }
    return true
}

func DFS(root int, visited []int, edges [][]int) bool {
    visited[root] = 1
    ret := true
    for _, link := range edges[root] {
        if visited[link] == 1 { return false }
        ret = ret && DFS(link, visited, edges)
        if ret == false { return false }
    }
    visited[root] = 2
    return true
}
