// monotonic stack
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    stack := []int{-1}
    var nums []int
    for idx := 0; head != nil; head, idx = head.Next, idx+1 {
        nums = append(nums, head.Val)
        for {
            if top := stack[len(stack)-1]; top != -1 && nums[top] < head.Val {
                nums[top] = head.Val
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        }
        stack = append(stack, idx)
    }
    for len(stack) > 1 {
        nums[stack[len(stack)-1]] = 0
        stack = stack[:len(stack)-1]
    }
    return nums
}
