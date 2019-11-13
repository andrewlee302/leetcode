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
import "math/rand"
func findKthLargest(nums []int, k int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        pivot := rand.Intn(r-l+1)+l
        nums[pivot], nums[r] = nums[r], nums[pivot]
        p, q := l, l
        for ; q < r; q++ {
            if nums[q] > nums[r] {
                nums[p], nums[q] = nums[q], nums[p]
                p++
            }
        }
        nums[p], nums[r] = nums[r], nums[p]
        if p == k - 1 {
            return nums[p]
        } else if p < k - 1 {
            l = p + 1
        } else {
            r = p - 1
        }
    }
    return -1
}
