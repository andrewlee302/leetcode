/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/list"
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stack1, stack2 := list.New(), list.New()
    for l1 != nil {
        stack1.PushBack(l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        stack2.PushBack(l2.Val)
        l2 = l2.Next
    }
    flag := 0
    var next *ListNode = nil
    for stack1.Len() != 0 || stack2.Len() != 0 {
        var v1, v2 int
        e1, e2 := stack1.Back(), stack2.Back()
        if e1 == nil { v1 = 0 } else { 
            stack1.Remove(e1)
            v1 = e1.Value.(int)
        }
        if e2 == nil { v2 = 0 } else {
            stack2.Remove(e2)
            v2 = e2.Value.(int)
        }
        sum := flag + v1 + v2
        flag = sum/10
        node := &ListNode{Val:sum%10, Next: next}
        next = node
    }
    if flag == 0 { return next } else { return &ListNode{Val:flag, Next: next} }
}
