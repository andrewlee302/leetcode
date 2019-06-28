// min-heap. Similar to 253. Meeting Rooms II.
import "container/heap"
import "sort"
func carPooling(trips [][]int, capacity int) bool {
    sort.Slice(trips, func (i, j int) bool { return trips[i][1] < trips[j][1] })
    h := &TripHeap{}
    for _, trip := range trips {
        for h.Len() != 0 && (*h)[0][2] <= trip[1] {
            stopTrip := heap.Pop(h).([]int)
            capacity += stopTrip[0]
        }
        capacity -= trip[0]
        if capacity < 0 { return false }
        heap.Push(h, trip)
    }
    return true
}

type TripHeap [][]int
func (h TripHeap) Len() int { return len(h) }
func (h TripHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h TripHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h *TripHeap) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *TripHeap) Pop() interface{} {
    v := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return v
}
