/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
func isCompleteTree(root *TreeNode) bool {
    if root == nil { return true }
    ok := true
    queue := list.New()
    queue.PushBack(root)
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            v := e.Value.(*TreeNode)
            if v.Left != nil {
                if !ok { return false }
                queue.PushBack(v.Left)
            } else {
                ok = false
            }
            if v.Right != nil {
                if !ok { return false }
                queue.PushBack(v.Right)
            } else {
                ok = false
            }
        }
    }
    return true
}
