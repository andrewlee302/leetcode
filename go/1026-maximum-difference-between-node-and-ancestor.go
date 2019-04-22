/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    var maxDiff int
    DFS(root, 100000, 0, &maxDiff)
    return maxDiff
}

func DFS(root *TreeNode, min, max int, maxDiff *int) {
    if root == nil { return }
    if root.Val < min { min = root.Val }
    if root.Val > max { max = root.Val }
    if *maxDiff < max - min { *maxDiff = max - min }
    DFS(root.Left, min, max, maxDiff)
    DFS(root.Right, min, max, maxDiff)
}
