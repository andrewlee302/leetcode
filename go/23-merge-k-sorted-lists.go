/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 { return nil }
    h := &ListHeap{}
    for _, list := range lists { if list != nil { heap.Push(h, list) } }
    dummy := &ListNode{Val:0}
    p := dummy
    for h.Len() != 0 {
        top := (*h)[0]
        p.Next = top
        p = top
        (*h)[0] = top.Next
        if (*h)[0] != nil { heap.Fix(h, 0) } else { heap.Remove(h, 0) }
    }
    p.Next = nil
    return dummy.Next
}

type ListHeap []*ListNode
func (h ListHeap) Len() int { return len(h) }
func (h ListHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h *ListHeap) Push(v interface{}) { *h = append(*h, v.(*ListNode)) }
func (h *ListHeap) Pop() interface{} {
    ret := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return ret
}
