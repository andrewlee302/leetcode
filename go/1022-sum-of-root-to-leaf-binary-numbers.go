/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const modulo int = 1000000007
func sumRootToLeaf(root *TreeNode) int {
    if root == nil { return 0 }
    var sum int
    DFS(root, 0, &sum)
    return sum
}

func DFS(root *TreeNode, remainder int, sum *int) {
    remainder = remainder*2 + root.Val
    remainder = remainder % modulo
    if root.Left != nil { DFS(root.Left, remainder, sum) }
    if root.Right != nil { DFS(root.Right, remainder, sum) }
    if root.Left == nil && root.Right == nil {
        *sum += remainder
        *sum = *sum % modulo
    }
}
