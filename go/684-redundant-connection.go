// DFS detect cycle in graph.
func findRedundantConnection(edges [][]int) []int {
    graph := make([][]int, 1001)
    for _, edge := range edges {
        visited := make([]bool, 1001)
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
        if DFS(graph, edge[0], -1, visited) { return edge }
    }
    return nil
}

// return whether is acyclic.
func DFS(graph [][]int, v, parent int, visited []bool) bool {
    visited[v] = true
    for _, link := range graph[v] {
        if !visited[link] {
            if DFS(graph, link, v, visited) { return true }
        } else if link != parent {
            return true
        }
    }
    return false
}
