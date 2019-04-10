/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 { return nil }
    root := &TreeNode{preorder[0], nil, nil}
    pivotIdx := 0
    for ; pivotIdx < len(inorder); pivotIdx++ {
        if inorder[pivotIdx] == preorder[0] { break }
    }
    // Number, left: pivot, right: len(preorder) - 1 - pivot
    root.Left = buildTree(preorder[1:pivotIdx+1], inorder[0:pivotIdx])
    root.Right = buildTree(preorder[pivotIdx+1:], inorder[pivotIdx+1:])
    return root
}
