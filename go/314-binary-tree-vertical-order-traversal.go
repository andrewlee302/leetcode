/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
type TreeNodeWrapper struct {
    node *TreeNode
    x int
}

func verticalOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    m := make(map[int][]int)
    l := list.New()
    l.PushBack(TreeNodeWrapper{node:root, x:0})
    min := int(^uint(0)>>1)
    max := ^min
    for l.Len() != 0 {
        e := l.Front()
        l.Remove(e)
        wrap := e.Value.(TreeNodeWrapper)
        if wrap.x < min { min = wrap.x }
        if wrap.x > max { max = wrap.x }
        m[wrap.x] = append(m[wrap.x], wrap.node.Val)
        if wrap.node.Left != nil { l.PushBack(TreeNodeWrapper{node:wrap.node.Left, x:wrap.x-1}) }
        if wrap.node.Right != nil { l.PushBack(TreeNodeWrapper{node:wrap.node.Right, x:wrap.x+1}) }
    }
    var ret [][]int
    for i := min; i <= max; i++ {
        if v, ok := m[i]; ok { ret = append(ret, v)}
    }
    return ret
}
