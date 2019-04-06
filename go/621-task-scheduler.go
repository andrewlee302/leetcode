import "container/heap"
func leastInterval(tasks []byte, n int) int {
    h := &IntHeap{} 
    heap.Init(h)
    counts := make(map[byte]int)
    for _, task := range tasks { counts[task]++ }
    for _, count := range counts { heap.Push(h, count) }
    time := 0
    for h.Len() != 0 {
        var restTaskCnts []int
        i := 1
        for ; h.Len() != 0 && i <= n+1; i++ {
            taskCnt := heap.Pop(h).(int)
            if taskCnt > 1 { restTaskCnts = append(restTaskCnts, taskCnt - 1) }
        }
        if h.Len() == 0 && len(restTaskCnts) == 0 {
            time += i - 1
        } else {
            time += n + 1
        }
        for _, count := range restTaskCnts { heap.Push(h, count) }
    }
    return time
}

type IntHeap []int
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Len() int { return len(h) }
func (h *IntHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntHeap) Pop() interface{} {
    v := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return v
}
