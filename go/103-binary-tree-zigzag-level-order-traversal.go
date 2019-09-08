/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    var ret [][]int
    var stack1, stack2 []*TreeNode
    stack1 = append(stack1, root)
    level := 0
    for len(stack1) != 0 {
        nums := make([]int, 0, len(stack1))
        for len(stack1) != 0 {
            node := stack1[len(stack1)-1]
            stack1 = stack1[:len(stack1)-1]
            nums = append(nums, node.Val)
            if level & 1 == 0 {
                if node.Left != nil { stack2 = append(stack2, node.Left) }
                if node.Right != nil { stack2 = append(stack2, node.Right) }
            } else {
                if node.Right != nil { stack2 = append(stack2, node.Right) }
                if node.Left != nil { stack2 = append(stack2, node.Left) }
            }
        }
        ret = append(ret, nums)
        stack1, stack2 = stack2, stack1
        level++
    }
    return ret
}
