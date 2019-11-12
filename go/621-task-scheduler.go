// Math: calculate idle slots.
func leastInterval(tasks []byte, n int) int {
    var cnts [26]int
    maxNum, maxTypeNum := 0, 0
    for _, task := range tasks {
        cnts[task-'A']++
        if num := cnts[task-'A']; num > maxNum {
            maxNum = cnts[task-'A']
            maxTypeNum = 1
        } else if num == maxNum {
            maxTypeNum++
        }
    }
    if maxTypeNum > n { return len(tasks) }
    otherTaskNum := len(tasks) - maxNum * maxTypeNum
    idleSlots := (maxNum - 1) * (n + 1 - maxTypeNum)
    if otherTaskNum >= idleSlots { return len(tasks) }
    return len(tasks) + (idleSlots - otherTaskNum)
}

// Greedy + max-heap.
import "container/heap"

func leastInterval(tasks []byte, n int) int {
    var cnts [26]int
    for _, task := range tasks { cnts[task-'A']++ }
    h := &IntHeap{}
    for _, cnt := range cnts { if cnt != 0 { heap.Push(h, cnt) } }
    time := 0
    taskCache := make([]int, 0, 26)
    for h.Len() > 0 {
        size := h.Len()
        i := 0
        for ; i < min(n + 1, size) ; i++ {
            if v := heap.Pop(h).(int) - 1; v > 0 {
                taskCache = append(taskCache, v)
            }
        }
        for _, n := range taskCache { heap.Push(h, n) }
        taskCache = taskCache[:0]
        if h.Len() == 0 { time += i } else { time += n + 1 }
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

func min(i, j int) int { if i < j { return i } else { return j } }
