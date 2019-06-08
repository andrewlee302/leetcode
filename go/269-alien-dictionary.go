import "container/list"
// topological sort
func alienOrder(words []string) string {
    graph := make(map[byte][]byte)
    indegree := make(map[byte]int)
    var order []byte
    // set indegree to zero if possible and compare the consecutive ones 
    // to get the connections.
    for i := 0; i < len(words); i++ {
        minL := len(words[i])
        if i < len(words) - 1 { minL = min(len(words[i]), len(words[i+1])) }
        isSet := false
        for k := 0; k < len(words[i]); k++ {
            currB := words[i][k]
            if indegree[currB] == 0 { indegree[currB] = 0 } // set if not exist
            if !isSet && i < len(words) - 1 && k < minL {
                nextB := words[i+1][k]
                if currB != nextB {
                    graph[currB] = append(graph[currB], nextB)
                    indegree[nextB]++
                    isSet = true
                }
            }
        }
    }
    // topological sort using indegrees and the graph.
    queue := list.New()
    for b, val := range indegree {
        if val == 0 { queue.PushBack(b) }
    }
    cnt := 0
    for queue.Len() != 0 {
        cnt++
        e := queue.Front()
        queue.Remove(e)
        currB := e.Value.(byte)
        order = append(order, currB)
        for _, nextB := range graph[currB] {
            indegree[nextB]--
            if indegree[nextB] == 0 { queue.PushBack(nextB) }
        }
    }
    if cnt != len(indegree) {
        return ""
    } else {
        return string(order)
    }
}

func min(a, b int) int { if a < b { return a } else { return b } }
func max(a, b int) int { if a < b { return b } else { return a } }
