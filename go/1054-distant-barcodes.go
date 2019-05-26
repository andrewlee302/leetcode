// greedy (fill the same number into every other slot)
import "sort"

func rearrangeBarcodes(barcodes []int) []int {
    // count frequency
    dict := make(map[int]int)
    for i := 0; i < len(barcodes); i++ {
        dict[barcodes[i]]++
    }
    var freqs [][]int
    for num, occur := range dict {
        freqs = append(freqs, []int{num, occur})
    }
    // in dscending order
    sort.Slice(freqs, func (i, j int) bool { return freqs[i][1] > freqs[j][1] })
    ret := make([]int, len(barcodes))
    slot := 0
    // fill the same numbers into every other slot.
    for _, freq := range freqs {
        for i := 0; i < freq[1]; i++ {
            // if exceed the length, go back to the slot with index 1.
            if slot >= len(barcodes) { slot = 1 }
            ret[slot] = freq[0]
            slot += 2
        }
    }
    return ret
}

// greedy (choose the most two frequent numbers into the slot consecutively)
import "container/heap"
func rearrangeBarcodes(barcodes []int) []int {
    // count frequency
    dict := make(map[int]int)
    for i := 0; i < len(barcodes); i++ {
        dict[barcodes[i]]++
    }
    // max-heap involving all counts
    h := &CountHeap{}
    heap.Init(h)
    for val, cnt := range dict {
        heap.Push(h, CodeCount{val, cnt})
    }
    ret := make([]int, 0, len(barcodes))
    // choose the maximum two counts and fill the two into the answer.
    for h.Len() >= 2 {
        cc1 := heap.Pop(h).(CodeCount)
        cc2 := heap.Pop(h).(CodeCount)
        // decrease
        cc1.cnt--
        cc2.cnt--
        ret = append(ret, cc1.val)
        ret = append(ret, cc2.val)
        // push back
        if cc1.cnt > 0 { heap.Push(h, cc1) }
        if cc2.cnt > 0 { heap.Push(h, cc2) }
    }
    // push the rest if exists, whose the count is definitely 1.
    if h.Len() == 1 {
        cc := heap.Pop(h).(CodeCount)
        ret = append(ret, cc.val) // cc.cnt must be 1.
    }
    return ret
}

type CodeCount struct {
    val, cnt int
}
type CountHeap []CodeCount
func (h CountHeap) Less (i ,j int) bool { return h[i].cnt > h[j].cnt }
func (h CountHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] }
func (h CountHeap) Len() int { return len(h) }
func (h *CountHeap) Push(x interface{}) { *h = append(*h, x.(CodeCount)) } // add x as element Len()
func (h *CountHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}
