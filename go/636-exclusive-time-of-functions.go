import "container/list"
import "strings"
import "strconv"
func exclusiveTime(n int, logs []string) []int {
    stack := list.New()
    ret := make([]int, n)
    for _, log := range logs {
        ss := strings.Split(log, ":")
        jobId, _ := strconv.Atoi(ss[0])
        time, _ := strconv.Atoi(ss[2])
        if ss[1][0] == 's' {
            stack.PushBack([]int{jobId, time})
        } else {
            e := stack.Back()
            tuple := e.Value.([]int)
            gap := 0
            for len(tuple) == 1 {
                stack.Remove(e)
                gap += tuple[0]
                e = stack.Back()
                tuple = e.Value.([]int)
            }
            stack.Remove(e)
            ret[jobId] += time - tuple[1] + 1 - gap
            stack.PushBack([]int{time - tuple[1] + 1})
        }
    }
    return ret
}
