/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var prev int
    isFirst := true
    return InOrderDFS(root, &prev, &isFirst)
}

func InOrderDFS(root *TreeNode, prev *int, isFirst *bool) bool {
    if root == nil { return true }
    if !InOrderDFS(root.Left, prev, isFirst) { return false }
    if !*isFirst && root.Val <= *prev { return false }
    *prev = root.Val
    *isFirst = false
    if !InOrderDFS(root.Right, prev, isFirst) { return false }
    return true
}
