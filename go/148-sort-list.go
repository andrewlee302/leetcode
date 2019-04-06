/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }
    var slowPrev *ListNode 
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slowPrev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    slowPrev.Next = nil
    left := sortList(head)
    right := sortList(slow)
    return merge(left, right) 
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
    dummy := new(ListNode)
    curr := dummy
    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val { 
            curr.Next = head1
            head1 = head1.Next
        } else {
            curr.Next = head2
            head2 = head2.Next
        }
        curr = curr.Next
    }
    if head1 != nil { curr.Next = head1 }
    if head2 != nil { curr.Next = head2 }
    return dummy.Next
}
