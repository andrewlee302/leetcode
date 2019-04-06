/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// iteration
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if p.Val > root.Val && q.Val > root.Val { 
            root = root.Right
        } else if p.Val < root.Val && q.Val < root.Val {
            root = root.Left  
        } else {
            return root
        } 
    }
    return nil
}

// recursion
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil { return nil }
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else {
        return root
    }
}
