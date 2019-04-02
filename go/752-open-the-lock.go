import "container/list"
import "strconv"
func openLock(deadends []string, target string) int {
    rotate := [][]byte{{'1','9'},{'0','2'},{'1','3'},{'2','4'},{'3','5'},{'4','6'},{'5','7'},{'6','8'},{'7','9'},{'8','0'},}
    if target == "0000" { return 0 }
    status := make([]int, 10000) // -1: dead, 0: unvisited, 1: visited
    for _, dead := range deadends { 
        d, _ := strconv.Atoi(dead)
        if d == 0 { return -1 } // Note "0000" could be in the deadends
        status[d] = -1
    }
    queue := list.New()
    queue.PushBack("0000")
    status[0] = 1
    level := 0
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            s := e.Value.(string)
            sb := []byte(s)
            for k := 0; k < 4; k++ {
                moveTo := rotate[int(s[k]-'0')]
                for l := 0; l < len(moveTo); l++ {
                    sb[k] = moveTo[l]
                    nextS := string(sb)
                    if target == nextS { return level + 1 }
                    nextSv, _ := strconv.Atoi(nextS)
                    if status[nextSv] == 0 {
                        status[nextSv] = 1
                        queue.PushBack(nextS)
                    }
                }
                sb[k] = s[k]
            }
        }
        level++
    }
    return -1
}
