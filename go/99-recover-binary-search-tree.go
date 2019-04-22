/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    var prev, mark1, mark2 *TreeNode = nil, nil, nil
    swapped := false
    DFS(root, &prev, &mark1, &mark2, &swapped)
    if !swapped {
        // if not found, just swap the adjecent ones.
        (*mark1).Val, (*mark2).Val = (*mark2).Val, (*mark1).Val
    }
}

func DFS(root *TreeNode, prev, mark1, mark2 **TreeNode, swapped *bool) {
    if root == nil { return }
    DFS(root.Left, prev, mark1, mark2, swapped)
    if *swapped { return }
    if *prev != nil && *mark1 == nil && (*prev).Val > root.Val {
        *mark1 = *prev
        *mark2 = root
    } else if *mark1 != nil && (*prev).Val > root.Val {
        // find the target and then swap
        (*mark1).Val, root.Val = root.Val, (*mark1).Val
        *swapped = true
        return
    }
    *prev = root
    DFS(root.Right, prev, mark1, mark2, swapped)
}
