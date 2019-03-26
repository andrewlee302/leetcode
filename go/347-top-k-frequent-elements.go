import "container/heap"

func topKFrequent(nums []int, k int) []int {
    if k <= 0 { return []int{} }
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ { m[nums[i]]++ }
    h := &KVHeap{}
    heap.Init(h)
    for key, value := range m {
        if len(*h) < k {
            heap.Push(h, KV{key, value})    
        } else if (*h)[0].value < value {
            // heap.Pop(h); heap.Push(h, KV{key, value})
            (*h)[0] = KV{key, value}
            heap.Fix(h, 0)
        }
    }
    ret := make([]int, len(*h))
    for i := 0; i < len(*h); i++ { ret[i] = (*h)[i].key }
    return ret
}

type KV struct {
    key, value int
}

type KVHeap []KV
func (h KVHeap) Less (i, j int) bool { return h[i].value < h[j].value }
func (h KVHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] }
func (h KVHeap) Len () int { return len(h) }
func (h *KVHeap) Push (x interface{}) { *h = append(*h, x.(KV)) }
func (h *KVHeap) Pop () interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return x
}
