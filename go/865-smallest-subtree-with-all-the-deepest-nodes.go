/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    _, ret := DFS(root)
    return ret
}

func DFS(root *TreeNode) (int, *TreeNode) {
    if root == nil { return 0, nil }
    LH, LRet := DFS(root.Left)
    RH, RRet := DFS(root.Right)
    if LH == RH { return LH+1, root }
    if LH > RH { return LH+1, LRet }
    return RH+1, RRet
}
