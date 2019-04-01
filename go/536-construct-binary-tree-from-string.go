/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
func str2tree(s string) *TreeNode {
    if len(s) == 0 { return nil }
    stack := list.New()
    for i := 0; i < len(s); {
        if s[i] == '(' {
            i++
        } else if s[i] == ')' {
            stack.Remove(stack.Back())
            i++
        } else {
            neg := false
            if s[i] == '-' { i, neg = i + 1, true }
            num := 0
            for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ { 
                num = num * 10 + int(s[i] - '0') 
            }
            if neg { num = -num }
            curr := &TreeNode{num, nil, nil}
            if stack.Len() != 0 { 
                p := stack.Back().Value.(*TreeNode)
                if p.Left != nil { p.Right = curr } else { p.Left = curr }
            }
            stack.PushBack(curr)
        }
    }
    return stack.Back().Value.(*TreeNode)
}
