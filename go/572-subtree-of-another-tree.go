// DFS.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
    return traversal(s, t)
}

func equals(x, y *TreeNode) bool {
    if x == nil && y == nil { return true }
    if x == nil || y == nil { return false }
    return x.Val == y.Val && equals(x.Left, y.Left) && equals(x.Right, y.Right)
}

func traversal(s, t *TreeNode) bool {
    return equals(s, t) || s != nil && (traversal(s.Left, t) || traversal(s.Right, t))
}
