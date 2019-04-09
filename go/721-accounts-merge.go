import "sort"
func accountsMerge(accounts [][]string) [][]string {
    graph := make(map[string][]string)
    emailNameMap := make(map[string]string)
    for _, account := range accounts {
        name := account[0]
        emailNameMap[account[1]] = name
        for _, email := range account[1:] { // Note: connect each other, even self. It doesn't matter.
            graph[email] = append(graph[email], account[1])
            graph[account[1]] = append(graph[account[1]], email)
            emailNameMap[email] = name
        }
    }
    var ret [][]string
    visited := make(map[string]bool)
    for email, _ := range graph {
        if !visited[email] {
            visited[email] = true
            set := []string{emailNameMap[email]}
            DFS(graph, visited, &set, email)
            ret = append(ret, set)
        }
    }
    for _, account := range ret {
        sort.Strings(account[1:])
    }
    return ret
}

func DFS(graph map[string][]string, visited map[string]bool, set *[]string, root string) {
    *set = append(*set, root)
    for _, email := range graph[root] {
        if !visited[email] {
            visited[email] = true
            DFS(graph, visited, set, email)
        }
    }
}
