/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil { return [][]int{} }
    var ret [][]int
    var queue []*TreeNode
    queue = append(queue, root)
    for len(queue) != 0 {
        size := len(queue)
        values := make([]int, 0, size)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            values = append(values, node.Val)
            if node.Left != nil { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        ret = append(ret, values)
    }
    return ret
}
