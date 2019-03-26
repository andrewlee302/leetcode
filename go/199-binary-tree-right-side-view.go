/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"

func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{} }
    var ret []int
    queue := list.New()
    queue.PushBack(root)
    for queue.Len() != 0 {
        size := queue.Len()
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            node := e.Value.(*TreeNode)
            if i == size - 1 { ret = append(ret, node.Val) }
            if node.Left != nil { queue.PushBack(node.Left) }
            if node.Right != nil { queue.PushBack(node.Right) }
        }
    }
    return ret
}
