// DFS
import "strconv"
const MAX_VISIT = 19900
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
    blockedMap := make(map[string]bool)
    for _, b := range blocked { blockedMap[cellStr(b[0], b[1])] = true }
    if blockedMap[cellStr(target[0], target[1])] { return false }
    return canAccess(blockedMap, source, target) && canAccess(blockedMap, target, source)
}

func canAccess(blockedMap map[string]bool, start, end []int) bool {
    visited := make(map[string]bool) 
    return dfs(blockedMap, visited, start, end)
}

func dfs(blockedMap, visited map[string]bool, start, end []int) bool {
    visited[cellStr(start[0], start[1])] = true
    if start[0] == end[0] && start[1] == end[1] || len(visited) > MAX_VISIT { return true }
    for _, move := range [][]int{{0,1},{1,0}, {-1,0}, {0,-1}} {
        nextX, nextY := start[0] + move[0], start[1] + move[1]
        nextKey := cellStr(nextX, nextY)
        if nextX >= 0 && nextX < 1e6 && nextY >= 0 && nextY < 1e6 && !blockedMap[nextKey] && !visited[nextKey] {
            if dfs(blockedMap, visited, []int{nextX, nextY}, end) { return true }
        }
    }
    return false
}

func cellStr(x, y int) string {
    return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
