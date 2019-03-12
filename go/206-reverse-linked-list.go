/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var last *ListNode = nil
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = last
        last = curr
        curr = next
    }
    return last
}
