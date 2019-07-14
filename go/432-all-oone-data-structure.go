// map + two heap(minheap & maxheap)
import "container/heap"
type AllOne struct {
    kv map[string]*Pair
    maxHeap *PairMaxHeap
    minHeap *PairMinHeap
}

/** Initialize your data structure here. */
func Constructor() AllOne {
    return AllOne{kv: make(map[string]*Pair), maxHeap: &PairMaxHeap{}, minHeap: &PairMinHeap{}}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
    if pair, ok := this.kv[key]; !ok {
        pair = &Pair{key: key, val: 1}
        this.kv[key] = pair
        heap.Push(this.maxHeap, pair)
        heap.Push(this.minHeap, pair)
    } else {
        this.kv[key].val++
        heap.Fix(this.maxHeap, pair.maxIdx)
        heap.Fix(this.minHeap, pair.minIdx)
    }
    fmt.Println("incr", this.kv)
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
    if pair, ok := this.kv[key]; ok {
        pair.val--
        if pair.val == 0 {
            delete(this.kv, key)
            saveMaxIdx, saveMinIdx := pair.maxIdx, pair.minIdx
            this.maxHeap.Swap(pair.maxIdx, this.maxHeap.Len()-1)
            this.maxHeap.Pop()
            this.minHeap.Swap(pair.minIdx, this.minHeap.Len()-1)
            this.minHeap.Pop()
            if saveMaxIdx < this.maxHeap.Len() { heap.Fix(this.maxHeap, saveMaxIdx) }
            if saveMinIdx < this.minHeap.Len() { heap.Fix(this.minHeap, saveMinIdx) } 
        } else {
            heap.Fix(this.maxHeap, pair.maxIdx)
            heap.Fix(this.minHeap, pair.minIdx)
        }
    } 
    fmt.Println("dec", this.kv)
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
    fmt.Println("GetMaxKey", this.kv)
    if this.maxHeap.Len() == 0 {
        return ""
    } else {
        return this.maxHeap.PairHeap[0].key
    }
    
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
    fmt.Println("GetMinKey", this.kv)
    if this.minHeap.Len() == 0 {
        return ""
    } else {
        return this.minHeap.PairHeap[0].key
    }
}

type Pair struct {
    key string
    val int
    maxIdx, minIdx int
}

type PairHeap []*Pair
type PairMaxHeap struct { PairHeap }
type PairMinHeap struct { PairHeap }

func (h PairHeap) Len() int { return len(h) }
func (h PairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h PairHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h *PairHeap) Push(v interface{}) { *h = append(*h, v.(*Pair)) }
func (h *PairHeap) Pop() interface{} {
    ret := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return ret
}

func (h PairMaxHeap) Swap(i, j int) {
    h.PairHeap.Swap(i, j)
    h.PairHeap[i].maxIdx, h.PairHeap[j].maxIdx = i, j
}
func (h PairMaxHeap) Less(i, j int) bool { return !h.PairHeap.Less(i, j) }
func (h *PairMaxHeap) Push(v interface{}) {
    h.PairHeap.Push(v)
    v.(*Pair).maxIdx = h.Len() - 1
}

func (h PairMinHeap) Swap(i, j int) {
    h.PairHeap.Swap(i, j)
    h.PairHeap[i].minIdx, h.PairHeap[j].minIdx = i, j
}
func (h *PairMinHeap) Push(v interface{}) {
    h.PairHeap.Push(v)
    v.(*Pair).minIdx = h.Len() - 1
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
