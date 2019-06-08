import "container/heap"
const INT_MAX = int(^uint(0)>>1)
const INT_MIN = ^INT_MAX
func smallestRange(nums [][]int) []int {
    minV, maxV := INT_MAX, INT_MIN
    h := &ListHeap{}
    heap.Init(h)
    for _, list := range nums {
        if len(list) != 0 {
            heap.Push(h, list)    
            minV = min(minV, list[0])
            maxV = max(maxV, list[0])
        }
    }
    minRange := []int{minV, maxV}
    for h.Len() == len(nums) { // there must exist k elements from k lists.
        if len((*h)[0]) > 1 {
            (*h)[0] = (*h)[0][1:]
            maxV = max(maxV, (*h)[0][0])
            heap.Fix(h, 0)
            minV = (*h)[0][0]
            if (maxV - minV) < (minRange[1] - minRange[0]) {
                minRange = []int{minV, maxV}
            }
        } else {
            heap.Pop(h)
        }
    }
    return minRange
}

type ListHeap [][]int
func (h ListHeap) Less (i ,j int) bool { return h[i][0] < h[j][0] }
func (h ListHeap) Swap (i, j int) { h[i], h[j] = h[j], h[i] } 
func (h ListHeap) Len() int { return len(h) }
func (h *ListHeap) Push(x interface{}) { *h = append(*h, x.([]int)) } // add x as element Len()
func (h *ListHeap) Pop() interface{} { // remove and return element Len() - 1.
    x := (*h)[len(*h)-1] // Note: (*h)[i]
    *h = (*h)[:len(*h)-1]
    return x
}

func min(a, b int) int { if a < b { return a } else { return b } }
func max(a, b int) int { if a < b { return b } else { return a } }
