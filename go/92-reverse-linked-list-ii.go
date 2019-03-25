/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// iterative version
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := new(ListNode)
    head, dummy.Next = dummy, head
    for i := 0; i < m-1; i++ { head = head.Next }
    var curr, prev *ListNode = head.Next, nil
    for i := 0; i < n - m + 1; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    head.Next.Next = curr
    head.Next = prev
    return dummy.Next
}

// recursive version
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := new(ListNode)
    head, dummy.Next = dummy, head
    for i := 0; i < m-1; i++ {
        head = head.Next
    }
    newHead, _ := reverseList(head.Next, n - m + 1)
    head.Next = newHead
    // if m == 1 { return head.Next } else { return dummy.Next }
    return dummy.Next
}

// return new head and the head of the rest
func reverseList(head *ListNode, cnt int) (*ListNode, *ListNode) {
    if cnt == 1 { return head, head.Next }
    newHead, restHead := reverseList(head.Next, cnt - 1)
    head.Next.Next = head
    head.Next = restHead
    return newHead, restHead
}
