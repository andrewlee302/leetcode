/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "sort"
type Element struct{ y, val int }
func verticalTraversal(root *TreeNode) [][]int {
    m := make(map[int][]Element)
    minX := int(^uint(0) >> 1)
    maxX := ^minX
    DFS(root, 0, 0, m, &minX, &maxX)
    var ret [][]int
    for k := minX; k <= maxX; k++ {
        if eles, ok := m[k]; ok {
            sort.Slice(eles, func (i, j int) bool { 
                if eles[i].y != eles[j].y { return eles[i].y > eles[j].y } else { return eles[i].val < eles[j].val }
            } )
            vals := make([]int, len(eles))
            for i, ele := range eles { vals[i] = ele.val }
            ret = append(ret, vals)
        }
    }
    return ret
}

func DFS(root *TreeNode, x, y int, m map[int][]Element, minX, maxX *int) {
    if root == nil { return }
    if x < *minX { *minX = x }
    if x > *maxX { *maxX = x }
    m[x] = append(m[x], Element{y, root.Val})
    DFS(root.Left, x-1, y-1, m, minX, maxX)
    DFS(root.Right, x+1, y-1, m, minX, maxX)
}
