/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    ret, _ := DFS(root)
    return ret
}

// LCA and the depth.
func DFS(root *TreeNode) (*TreeNode, int) {
    if root == nil { return nil, -1 }
    lLCA, lDepth := DFS(root.Left)
    rLCA, rDepth := DFS(root.Right)
    if lDepth == rDepth {
        return root, lDepth+1
    } else if lDepth > rDepth {
        return lLCA, lDepth+1
    } else {
        return rLCA, rDepth+1
    }
}
