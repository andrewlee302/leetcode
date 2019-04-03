import "container/heap"
type WordCnt struct {
    word string
    cnt int
}

type WordCntHeap []WordCnt

func (h WordCntHeap) Less(i, j int) bool { 
    if h[i].cnt != h[j].cnt { return h[i].cnt < h[j].cnt } else { return h[i].word > h[j].word }
}
func (h WordCntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h WordCntHeap) Len() int { return len(h) }
func (h *WordCntHeap) Push(v interface{}) { *h = append(*h, v.(WordCnt)) }
func (h *WordCntHeap) Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return x
}

func topKFrequent(words []string, k int) []string {
    h := &WordCntHeap{}
    m := make(map[string]int)
    for _, w := range words { m[w]++ }
    heap.Init(h)
    for word, cnt := range m {
        if h.Len() < k { heap.Push(h, WordCnt{word, cnt}) } else {
            if (*h)[0].cnt < cnt || (*h)[0].cnt == cnt && word < (*h)[0].word {
                (*h)[0] = WordCnt{word, cnt}
                heap.Fix(h, 0)
            }
        }
    }
    ret := make([]string, k)
    for i := k-1; i >= 0; i-- { ret[i] = heap.Pop(h).(WordCnt).word }
    return ret
}
