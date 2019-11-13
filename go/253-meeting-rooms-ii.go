// Slots.
import "sort"
type TimeFlag struct {
    time int
    flag int // +1, -1
}
func minMeetingRooms(intervals [][]int) int {
    var flags []TimeFlag
    for _, itv := range intervals {
        flags = append(flags, TimeFlag{itv[0], 1})
        flags = append(flags, TimeFlag{itv[1], -1})
    }
    sort.Slice(flags, func (i, j int) bool {
        if flags[i].time != flags[j].time {
            return flags[i].time < flags[j].time
        } else {
            return flags[i].flag < flags[j].flag
        }
    })
    occupy := 0
    ret := 0
    for _, item := range flags {
        occupy += item.flag
        ret = max(ret, occupy)
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }

// Min-heap.
import "sort"
import "container/heap"
func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    h := &IntHeap{}
    ret := 0
    for _, interval := range intervals {
        // release rooms
        for h.Len() != 0 && (*h)[0] <= interval[0] {
            heap.Pop(h)
        }
        heap.Push(h, interval[1])
        if h.Len() > ret { ret = h.Len() }
    }
    return ret
}

// minimum heap. If we want maximum heap, change the Less function.
type IntHeap []int
func (h IntHeap) Less (i ,j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] } 
func (h IntHeap) Len() int { return len(h) }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) } // add x as element Len()
func (h *IntHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}
