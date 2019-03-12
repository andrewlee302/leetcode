/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil {return true}
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    var prev *ListNode = nil
    curr := head
    for curr != slow {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    curr = prev
    if fast != nil {slow = slow.Next}
    for curr != nil {
        if curr.Val != slow.Val {
            return false
        }
        curr = curr.Next
        slow = slow.Next
    }
    return true
}
