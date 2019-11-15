// Binary search.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil { return 0 }
    depth := 0
    for p := root; p != nil; p = p.Left { depth++ } // root: 1.
    // maxN := 2^depth - 1, last_level_maxN = 2^(depth-1)
    l, r := 1, 1 << uint(depth-1) + 1 // Note, r value.
    // find samllest position where there is no leaf node on the last level.
    for l < r {
        mid := l + (r-l)/2
        ll, rr := 1, 1 << uint(depth-1)
        p := root
        for d := 1; d < depth; d++ {
            pivot := ll + (rr-ll)/2
            if mid <= pivot {
                p = p.Left
                rr = pivot
            } else {
                p = p.Right
                ll = pivot + 1
            }
        }
        if p == nil { // no leaf node.
            r = mid
        } else {
            l = mid + 1
        }
    }
    return 1 << uint(depth-1) - 1 + l - 1
}
