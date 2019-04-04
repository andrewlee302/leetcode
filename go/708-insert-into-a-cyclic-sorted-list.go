/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
func flatten(root *TreeNode)  {
    if root == nil { return }
    stack := list.New()
    stack.PushBack(root)
    prev := new(TreeNode)
    for stack.Len() != 0 {
        e := stack.Back()
        stack.Remove(e)
        node := e.Value.(*TreeNode)
        left, right := node.Left, node.Right
        prev.Right = node
        prev.Left = nil
        prev = node
        if right != nil { stack.PushBack(right) }
        if left != nil { stack.PushBack(left) }
    }
    root = prev.Right
}


