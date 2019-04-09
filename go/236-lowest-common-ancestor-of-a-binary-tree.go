/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    _, ret := DFS(root, p, q)
    return ret
}

func DFS(root, p, q *TreeNode) (int, *TreeNode) {
    if root == nil { return 0, nil }
    var match = 0
    if root.Val == p.Val || root.Val == q.Val { match = 1 }
    leftMatch, leftLCA := DFS(root.Left, p, q)
    rightMatch, rightLCA := DFS(root.Right, p, q)
    if leftMatch == 2 { return 2, leftLCA }
    if rightMatch == 2 { return 2, rightLCA }
    match += leftMatch + rightMatch
    if match == 2 { return match, root } else { return match, nil }
}
