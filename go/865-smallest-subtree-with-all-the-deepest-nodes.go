/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    node, _ := dfs(root, 0)
    return node
}

// root is not nil
func dfs(root *TreeNode, depth int) (*TreeNode, int) {
    if root.Left == nil && root.Right == nil { return root, depth }
    var l, r *TreeNode
    var lDepth, rDepth int
    if root.Left != nil {  l, lDepth = dfs(root.Left, depth+1) }
    if root.Right != nil {  r, rDepth = dfs(root.Right, depth+1) }
    if l != nil && r != nil { 
        if lDepth < rDepth {
            return r, rDepth
        } else if lDepth > rDepth {
            return l, lDepth
        } else {
            return root, lDepth
        }
    }
    if l != nil { return l, lDepth } else { return r, rDepth }
}
