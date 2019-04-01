import "sort"
func findItinerary(tickets [][]string) []string {
    links := make(map[string][]string)
    for i := 0; i < len(tickets); i++ {
        from, to := tickets[i][0], tickets[i][1]
        links[from] = append(links[from], to)
    }
    for _, tos := range links {
        sort.Strings(tos)
    }
    route := []string{"JFK"}
    DFS("JFK", links, &route, len(tickets)+1)
    return route
}

func DFS(root string, links map[string][]string, route *[]string, ticketNum int) bool {
    if len(*route) == ticketNum { return true }
    tos := links[root]
    for i := 0; i < len(tos); i++ {
        save := tos[i]
        if save != "0" {
            tos[i] = "0"
            *route = append(*route, save)
            if DFS(save, links, route, ticketNum) { return true }
            *route = (*route)[:len(*route)-1]
            tos[i] = save
        } 
    }
    return false
}
