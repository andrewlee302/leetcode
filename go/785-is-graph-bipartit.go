// BFS + colors.
import "container/list"
func isBipartite(graph [][]int) bool {
    colors := make(map[int]int) // 0, 1, 2
    for i := 0; i < len(graph); i++ {
        if colors[i] == 0 {
            queue := list.New()
            queue.PushBack(i)
            colors[i] = 1
            for queue.Len() != 0 {
                e := queue.Front()
                queue.Remove(e)
                v := e.Value.(int)
                nextColor := 3 - colors[v]
                for _, next := range graph[v] {
                    if colors[next] == 0 {
                        colors[next] = nextColor
                        queue.PushBack(next)
                    } else if colors[next] != nextColor { 
                        return false
                    }
                }
            }
        }
    }
    return true
}

// DFS (Iteration) + colors.
func isBipartite(graph [][]int) bool {
    colors := make(map[int]int) // 0, 1, 2
    for i := 0; i < len(graph); i++ {
        if colors[i] == 0 {
            colors[i] = 1
            var stack []int
            stack = append(stack, i)
            for len(stack) != 0 {
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                nextColor := 3 - colors[top]
                for _, next := range graph[top] {
                    if colors[next] == 0 {
                        colors[next] = nextColor
                        stack = append(stack, next)
                    } else if colors[next] != nextColor {
                        return false
                    }
                }
            }
        }
    }
    return true
}
