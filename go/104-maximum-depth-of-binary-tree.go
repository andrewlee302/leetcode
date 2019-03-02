/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    mdepth := 0 
    DFS(root, 0, &mdepth)
    return mdepth
}

// depth means the depth before visiting the subtree's root node
// root should be guaranteed not to be nil.
func DFS(root *TreeNode, depth int, mdepth *int) {
    depth ++
    if root.Left == nil && root.Right == nil && depth > *mdepth {
        *mdepth = depth
    }
    if root.Left != nil {
        DFS(root.Left, depth, mdepth)
    }
    if root.Right != nil {
        DFS(root.Right, depth, mdepth)
    }
    return
}
