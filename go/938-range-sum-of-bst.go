func rangeSumBST(root *TreeNode, L, R int) int {
    if root == nil || L > R { return 0 }
    ret := rangeSumBST(root.Left, L, min(R, root.Val-1)) + rangeSumBST(root.Right, max(L, root.Val+1), R)
    if root.Val >= L && root.Val <= R { ret += root.Val }
    return ret
}

func min(a, b int) int { if a < b { return a } else { return b } }
func max(a, b int) int { if a > b { return a } else { return b } }
