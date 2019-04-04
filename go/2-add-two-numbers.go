/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    flag := 0
    head := &ListNode{0, nil}
    prev := head
    for l1 != nil || l2 != nil {
        v1, v2 := 0, 0
        if l1 != nil { 
            v1 = l1.Val 
            l1 = l1.Next
        }
        if l2 != nil { 
            v2 = l2.Val 
            l2 = l2.Next
        }
        sum := v1 + v2 + flag
        flag = sum/10
        node := &ListNode{sum%10, nil}
        prev.Next = node
        prev = node
    }
    if flag != 0 {
        node := &ListNode{flag, nil}
        prev.Next = node
    }
    return head.Next
}
