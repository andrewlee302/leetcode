/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"
import "fmt"
func mergeKLists(lists []*ListNode) *ListNode {
    var head, tail *ListNode
    h := &ListHeap{}
	heap.Init(h)
    for _, list := range lists {
        if list != nil {
            heap.Push(h, list)  
        }
    }
    for h.Len() > 0 {
        minList := heap.Pop(h).(*ListNode)
        if head == nil {
            head = minList
        } else {
            tail.Next = minList
        }
        tail = minList
        if minList.Next != nil {
            heap.Push(h, minList.Next)
        }   
    }
    return head
}

type ListHeap []*ListNode

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

