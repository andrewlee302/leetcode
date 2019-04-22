/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var prev *TreeNode = nil
    isValid := true
    DFS(root, &prev, &isValid)
    return isValid
}

func DFS(root *TreeNode, prev **TreeNode, isValid *bool) {
    if root == nil { return }
    DFS(root.Left, prev, isValid)
    if !*isValid { return }
    if *prev != nil && (*prev).Val >= root.Val {
        *isValid = false
        return
    }
    *prev = root
    DFS(root.Right, prev, isValid)
}
