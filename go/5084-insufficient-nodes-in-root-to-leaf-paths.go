/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// DFS recursion + parameters.
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    del := DFS(root, 0, limit)
    if del { return nil }
    return root
}

func DFS(root *TreeNode, sum, limit int) bool {
    if root == nil { return true }
    sum = sum + root.Val
    if root.Left == nil && root.Right == nil {
        if sum < limit {
            return true
        } else {
            return false
        }
    }
    lDel := DFS(root.Left, sum, limit)
    if lDel { root.Left = nil }
    rDel := DFS(root.Right, sum, limit)
    if rDel { root.Right = nil }
    return lDel && rDel
}
