/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    diameter := 0
    depth(root, &diameter)
    return diameter
}

func depth(node *TreeNode, diameter *int) int {
    if node == nil {
        return 1
    }
    leftDepth := depth(node.Left, diameter)
    rightDepth := depth(node.Right, diameter)
    var d int
    if rightDepth > leftDepth {
        d = rightDepth + 1
    } else {
        d = leftDepth + 1
    }
    if diameterTry := leftDepth + rightDepth - 2; diameterTry > *diameter {
        *diameter = diameterTry
    }
    return d
}
