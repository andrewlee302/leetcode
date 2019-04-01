/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    return helper(&preorder, inorder)
}

func helper(preorder *[]int, inorder []int) *TreeNode {
    if len(*preorder) == 0 || len(inorder) == 0 { return nil }
    root := &TreeNode{(*preorder)[0], nil, nil}
    i := 0
    for ; i < len(inorder) && inorder[i] != (*preorder)[0]; i++ { }
    *preorder = (*preorder)[1:]
    root.Left = helper(preorder, inorder[:i]) 
    root.Right = helper(preorder, inorder[i+1:]) 
    return root
}
