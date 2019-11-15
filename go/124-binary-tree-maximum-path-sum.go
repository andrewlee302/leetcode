/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxV := root.Val
    maxPathStartsWith(root, &maxV)
    return maxV
}

// return: maxPath starting with root
func maxPathStartsWith(root *TreeNode, maxV *int) int {
    if root == nil { return 0 }
    ls, rs := maxPathStartsWith(root.Left, maxV), maxPathStartsWith(root.Right, maxV)
    startLen := max(max(ls, rs), 0) + root.Val
    throughLen := max(ls, 0) + max(rs, 0) + root.Val // including starting with root.
    *maxV = max(throughLen, *maxV)
    return startLen
}

func max(i, j int) int { if i > j { return i } else { return j } }
