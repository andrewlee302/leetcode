/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // In-order recursion. 
func kthSmallest(root *TreeNode, k int) int {
    target, cnt := 0, 0
    DFS(root, &cnt, &target, k)
    return target
}

func DFS(root *TreeNode, cnt, target *int, k int) bool {
    if root == nil { return false }
    found := DFS(root.Left, cnt, target, k)
    if found { return true } 
    *cnt++
    if *cnt == k { 
        *target = root.Val
        return true 
    }
    found = DFS(root.Right, cnt, target, k)
    return found
}

// In-order iteration.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var stack []*TreeNode
    for node := root; node != nil; node = node.Left {
        stack = append(stack, node)
    }
    for cnt := 1; len(stack) != 0; cnt++ {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if cnt == k { return node.Val }
        for node = node.Right; node != nil; node = node.Left {
            stack = append(stack, node)
        }
    }
    return -1
}
