/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k == 1 {
        return head
    }
    var prev *ListNode = nil
    curr := head
    i := 0
    for ; i < k && curr != nil; i++ {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    // reverse again
    if i < k {
        return reverseKGroup(prev, i)
    }
    head.Next = reverseKGroup(curr, k)
    return prev
}
