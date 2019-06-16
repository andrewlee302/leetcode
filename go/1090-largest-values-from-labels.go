// Sort all and greedy.
import "sort"
func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
    items := make([]Item, len(values))
    for i := 0; i < len(values); i++ {
        items[i] = Item{values[i], labels[i]}
    }
    sort.Slice(items, func (i, j int) bool { return items[i].val > items[j].val })
    used := make(map[int]int)
    cnt, sum := 0, 0
    for i := 0; i < len(items) && cnt < num_wanted; i++ {
        if used[items[i].label] < use_limit {
            used[items[i].label]++
            sum += items[i].val
            cnt++
        }
    }
    return sum
}

type Item struct {
    val, label int
}

// Heap of heaps.
import "container/heap"
func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
    hh := &IntHeapLabelHeap{}
    m := make(map[int]*IntHeap)
    for i := 0; i < len(values); i++ {
        if m[labels[i]] == nil {
            m[labels[i]] = &IntHeap{}
        }
        heap.Push(m[labels[i]], values[i])
    }
    for label, h := range m { heap.Push(hh, IntHeapLabel{h, label}) }
    used := make(map[int]int)
    sum := 0
    for i := 0; i < num_wanted && hh.Len() != 0; i++ {
        h := heap.Pop(hh).(IntHeapLabel)
        sum += heap.Pop(h.nums).(int)
        used[h.label]++
        if h.nums.Len() != 0 && used[h.label] < use_limit {
            heap.Push(hh, h)
        }
    }
    return sum
}

// max-heap
type IntHeap []int
func (h IntHeap) Less (i ,j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Len() int { return len(h) }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) } // add x as element Len()
func (h *IntHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}

type IntHeapLabel struct {
    nums *IntHeap
    label int
}

// max-heap
type IntHeapLabelHeap []IntHeapLabel
func (h IntHeapLabelHeap) Less (i ,j int) bool { return (*(h[i].nums))[0] > (*(h[j].nums))[0] }
func (h IntHeapLabelHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] }
func (h IntHeapLabelHeap) Len() int { return len(h) }
func (h *IntHeapLabelHeap) Push(x interface{}) { *h = append(*h, x.(IntHeapLabel)) } // add x as element Len()
func (h *IntHeapLabelHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}
