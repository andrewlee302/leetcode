/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"
func closestValue(root *TreeNode, target float64) int {
    var closestVal int
    var diff = math.MaxFloat64
    DFS(root, target, &closestVal, &diff)
    return closestVal
}

func DFS(node *TreeNode, target float64, closestVal *int, diff *float64) {
    if node == nil {
        return
    }
    currDiff := math.Abs(float64(node.Val) - target)
    if  currDiff < *diff {
        *closestVal = node.Val
        *diff = currDiff
    }
    DFS(node.Left, target, closestVal, diff)
    DFS(node.Right, target, closestVal, diff)
}
