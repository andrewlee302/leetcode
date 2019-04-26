// monotonic queue
import "container/list"
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) < k { return []int{} }
    var ret []int
    queue := list.New()
    for i := 0; i < len(nums); i++ {
        for queue.Len() != 0 {
            e := queue.Back()
            if e.Value.(int) < nums[i] {
                queue.Remove(e)
            } else { break }
        }
        queue.PushBack(nums[i])
        // We can have a window.
        if i >= k - 1 {
            e := queue.Front()
            currWindowMaximum := e.Value.(int)
            ret = append(ret, currWindowMaximum)
            if nums[i-k+1] == currWindowMaximum {
                queue.Remove(e)
            }
        }
    }
    return ret
}

// max-queue with position tag.
import "container/heap"
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || len(nums) < k { return []int{} }
    ret := make([]int, len(nums)-k+1)
    window := make([]*Node, k) // circular array
    windowHead := 0
    h := &NodeHeap{}
    heap.Init(h)
    for i := 0; i < k; i++ {
        node := &Node{nums[i], i}
        window[i]= node
        heap.Push(h, node)
    }
    ret[0] = (*h)[0].val
    for i := 0; i <= len(nums) - k - 1; i++ {
        headNode := window[windowHead]
        heap.Remove(h, headNode.pos)
        tailNode := &Node{nums[i+k],len(*h)}
        window[windowHead] = tailNode
        heap.Push(h, tailNode)
        ret[i+1] = (*h)[0].val
        windowHead = (windowHead+1)%k
    }
    return ret
}

type Node struct {
    val, pos int
}
type NodeHeap []*Node
func (h NodeHeap) Swap (i ,j int) {
    h[i], h[j] = h[j], h[i]
    h[i].pos, h[j].pos = i, j
}
func (h NodeHeap) Less (i, j int) bool { return h[i].val > h[j].val } // max-heap
func (h NodeHeap) Len() int { return len(h) }
func (h *NodeHeap) Push(v interface{}) { *h = append(*h, v.(*Node)) }
func (h *NodeHeap) Pop() interface{} {
    ret := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return ret
}
