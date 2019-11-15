// Quickselect.
import "math/rand"
type Occur struct {
    num, freq int
}
func topKFrequent(nums []int, k int) []int {
    cnts := make(map[int]int)
    for _, num := range nums {
        cnts[num]++
    }
    arr := make([]Occur, len(cnts))
    for num, cnt := range cnts {
        arr = append(arr, Occur{num, cnt})
    }
    l, r := 0, len(arr) - 1
    for l <= r { // l < r is also correct
        pivot := rand.Intn(r-l+1) + l
        arr[pivot], arr[r] = arr[r], arr[pivot]
        p, q := l, l
        for ; q < r; q++ {
            if arr[q].freq > arr[r].freq {
                arr[p], arr[q] = arr[q], arr[p]
                p++
            }
        }
        arr[p], arr[r] = arr[r], arr[p]
        if p == k-1 || p == k {
            break
        } else if p > k {
            r = p - 1
        } else {
            l = p + 1
        }
    }
    ret := make([]int, k)
    for i := 0; i < k; i++ {
        ret[i] = arr[i].num
    }
    return ret
}

// Bucket sort.
func topKFrequent(nums []int, k int) []int {
    buckets := make([][]int, len(nums)+1)
    cnts := make(map[int]int)
    for _, num := range nums {
        cnts[num]++
    }
    for num, cnt := range cnts {
        buckets[cnt] = append(buckets[cnt], num)
    }
    var ret []int
    for i := len(nums); k > 0 && i >= 0; i-- {
        if l := len(buckets[i]); l > 0 {
            ret = append(ret, buckets[i][:min(k, l)]...)
            k -= min(k, l)
        }
    }
    return ret
}

func min(i, j int) int { if i < j { return i } else { return j } }

// Min-heap.
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
