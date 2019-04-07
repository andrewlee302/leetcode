import "container/list"

func isBipartite(graph [][]int) bool {
    colors := make([]int, len(graph)) // 0: unvisited, 1: blue, 2: red
    for i := 0; i < len(graph); i++ {
        if colors[i] != 0 { continue }
        stack := list.New()
        stack.PushBack(i)
        colors[i] = 1 // it's the same to set it to 2.
        for stack.Len() != 0 {
            e := stack.Back()
            top := e.Value.(int)
            stack.Remove(e)
            for _, v := range graph[top] {
                expectedColor := 3 - colors[top]
                if colors[v] == 0 { 
                    colors[v] = expectedColor
                    stack.PushBack(v) 
                } else if colors[v] != expectedColor { return false }
            }
        }
    }
    return true
}

// recursion
func isBipartite(graph [][]int) bool {
    colors := make([]int, len(graph)) // 0: unvisited, 1: blue, 2: red
    for i := 0; i < len(graph); i++ {
        if colors[i] == 0 {
            colors[i] = 1
            if !DFS(i, 1, graph, colors) { return false }
        }
    }
    return true
}

// setIdx describes the root's set.
func DFS(root int, color int, graph [][]int, colors []int) bool {
    colors[root] = color
    for _, node := range graph[root] {
        if colors[node] == 0 {
            if !DFS(node, 2-color+1, graph, colors) { return false }
        } else {
            if colors[node] != 2-color+1 { return false }
        }
    }
    return true
}
