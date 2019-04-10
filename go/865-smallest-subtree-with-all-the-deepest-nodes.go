/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    ret, _ := recursion(root)
    return ret
}

func recursion(root *TreeNode) (*TreeNode, int) {
    if root == nil { return nil, 0 }
    leftRet, leftDepth := recursion(root.Left)
    rightRet, rightDepth := recursion(root.Right)
    var depth int
    var ret *TreeNode
    if leftDepth == rightDepth {
        depth, ret = leftDepth + 1, root
    } else if leftDepth > rightDepth {
        depth, ret = leftDepth + 1, leftRet
    } else {
        depth, ret = rightDepth + 1, rightRet
    }
    return ret, depth
}

// depth = max(leftDepth, rightDepth) + 1
// smallestSubTree: 
// if leftDepth == rightDepth -> root
// else if leftDepth > rightDepth -> leftRet
// else -> rightRet
