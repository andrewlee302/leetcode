/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diameter, _ := DFS(root)
    return diameter
}
// diameter(root) = max(diameter(root.Left), diameter(root.Right), leftDepth+rightDepth+1)
// return the diameter and depth
func DFS(root *TreeNode) (int, int) {
    if root == nil { return 0, 0 }
    leftDiameter, leftDepth := DFS(root.Left)
    rightDiameter, rightDepth := DFS(root.Right)
    passRootPath := leftDepth + rightDepth
    diameter := passRootPath
    if leftDiameter > diameter { diameter = leftDiameter }
    if rightDiameter > diameter { diameter = rightDiameter }
    depth := 1 + leftDepth
    if 1 + rightDepth > depth { depth = 1 + rightDepth }
    return diameter, depth
}
