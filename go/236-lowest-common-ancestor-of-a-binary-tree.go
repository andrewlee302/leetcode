/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var ret *TreeNode
    DFS(root, p, q, &ret)
    return ret
}

func DFS(root, p, q *TreeNode, ret **TreeNode) int {
    if root == nil { return 0 }
    if *ret != nil { return 0 } // prune
    match := 0
    if root.Val == p.Val || root.Val == q.Val { match = 1 }
    match += DFS(root.Left, p, q, ret)
    match += DFS(root.Right, p, q, ret)
    if match >= 2 && *ret == nil { *ret = root }
    return match
}
