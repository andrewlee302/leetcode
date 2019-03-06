/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "container/heap"
import "sort"
func minMeetingRooms(intervals []Interval) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Sort(Intervals(intervals))
    h := &IntHeap{intervals[0].End}
    heap.Init(h)
    for _, itv := range intervals[1:] {
        earliestEnd := heap.Pop(h).(int)
        if earliestEnd <= itv.Start {
            heap.Push(h, itv.End)
        } else {
            heap.Push(h, earliestEnd)
            heap.Push(h, itv.End)
        }
    } 
    return len(*h)
}

type IntHeap []int
func (h IntHeap) Less (i, j int) bool { return h[i] < h[j]}
func (h IntHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i]}
func (h IntHeap) Len() int { return len(h)}
func (h *IntHeap) Push(v interface{}) {
    *h = append(*h, v.(int))
}
func (h *IntHeap) Pop() interface{} {
    l := len(*h)
    v := (*h)[l-1]
    *h = (*h)[:l-1]
    return v
}

type Intervals []Interval
func (itl Intervals) Less (i, j int) bool { return itl[i].Start < itl[j].Start}
func (itl Intervals) Swap (i, j int) { itl[i], itl[j] = itl[j], itl[i]}
func (itl Intervals) Len() int { return len(itl)}


