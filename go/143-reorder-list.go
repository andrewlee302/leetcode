/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil { return }
    prev, slow, fast := head, head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    var firstTail, secondHead *ListNode
    if fast == nil { 
        firstTail = prev 
        secondHead = slow
    } else { 
        firstTail = slow 
        secondHead = slow.Next
    }
    firstTail.Next = nil
    prev = nil
    for secondHead != nil {
        temp := secondHead.Next
        secondHead.Next = prev
        prev = secondHead
        secondHead = temp
    }
    secondHead = prev
    yummy := &ListNode{}
    for head != nil && secondHead != nil {
        tempHead, tempSecondHead := head.Next, secondHead.Next
        yummy.Next = head
        yummy.Next.Next = secondHead
        yummy = secondHead
        head, secondHead = tempHead, tempSecondHead
    }
    yummy.Next = head
}

