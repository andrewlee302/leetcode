// similar to leetcode 295.
import "container/heap"
import "container/list"
func medianSlidingWindow(nums []int, k int) []float64 {
    records := list.New()
    // left is one longer than or equal to the right queue.
    left := &MaxHeap{}
    right := &MinHeap{}
    ret := make([]float64, len(nums) - k + 1)
    // invariance: left.Len() = right.Len() + 1 or left.Len() == right.Len()
    for i := 0; i < len(nums); i++ {
        rec := &Rec{v:nums[i]}
        records.PushBack(rec)
        
        if left.Len() == right.Len() {
            rec.left = false
            heap.Push(right, rec)
            rightToLeft(left, right)
        } else {
            rec.left = true
            heap.Push(left, rec)
            leftToRight(left, right)
        }
        // here, we can guarantee left.Len() mustn't be less than right.Len().
        if i >= k {
            e := records.Front()
            records.Remove(e)
            delRec := e.Value.(*Rec)
            fmt.Println("remove", *delRec)
            if delRec.left {
                heap.Remove(left, delRec.idx)
            } else {
                heap.Remove(right, delRec.idx)
            }
            
            if left.Len() == right.Len() + 2 {
                leftToRight(left, right)
            } else if left.Len() == right.Len() - 1 {
                rightToLeft(left, right)
            } else {
                // left.Len() == right.Len() || left.Len() == right.Len() + 1
            }
        } 
        
        if i >= k - 1 {
            if k % 2 == 0 {
                ret[i-k+1] = float64(left.RecHeap[0].v + right.RecHeap[0].v) / 2.0
            } else {
                ret[i-k+1] = float64(left.RecHeap[0].v)
            }
        }
    }
    return ret   
}
// move the maxium of left to the right queue.
func leftToRight(left *MaxHeap, right *MinHeap) {
    pRec := heap.Pop(left).(*Rec)
    pRec.left = false
    heap.Push(right, pRec)
}
// move the minimal of right to the left queue.
func rightToLeft(left *MaxHeap, right *MinHeap) {
    pRec := heap.Pop(right).(*Rec)
    pRec.left = true
    heap.Push(left, pRec)
}

type Rec struct {
    v int
    idx int // idx in the heap, used in heap.Remove(h, idx)
    left bool  // is in the left or the right.
}
type RecHeap []*Rec
type MaxHeap struct { RecHeap }
type MinHeap struct { RecHeap }
func (h MinHeap) Less (i ,j int) bool { return h.RecHeap[i].v < h.RecHeap[j].v }
func (h MaxHeap) Less (i ,j int) bool { return h.RecHeap[i].v > h.RecHeap[j].v }
func (h RecHeap) Swap (i, j int) { 
    h[i], h[j] = h[j], h[i] 
    h[i].idx, h[j].idx = i, j
} 
func (h RecHeap) Len() int { return len(h) }
func (h *RecHeap) Push(x interface{}) { // add x as element Len()
    rec := x.(*Rec)
    rec.idx = len(*h)
    *h = append(*h, rec) 
} 
func (h *RecHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}
