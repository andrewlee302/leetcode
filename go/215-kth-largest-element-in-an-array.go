// O(N*logK), O(K).
import "container/heap"
func findKthLargest(nums []int, k int) int {
    h := &IntHeap{}
    for _, num := range nums {
        if h.Len() < k {
            heap.Push(h, num)
        } else if (*h)[0] < num {
            (*h)[0] = num
            heap.Fix(h, 0)
        }
    }
    return (*h)[0]
}

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

// O(N), O(N).
func findKthLargest(nums []int, k int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        pivot := nums[r]
        i, p := l, l
        for ; i < r; i++ {
            if nums[i] > pivot { 
                nums[i], nums[p] = nums[p], nums[i] 
                p++
            }
        }
        nums[r], nums[p] = nums[p], nums[r] // p is the pivotIndex
        if p - l >= k {
            r = p - 1
        } else if p - l + 1 == k {
            return nums[p]
        } else {
            k = k - (p - l + 1)
            l = p + 1
        }
    }
    return -1
}
