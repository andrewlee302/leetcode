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

