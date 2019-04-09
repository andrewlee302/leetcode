/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func str2tree(s string) *TreeNode {
    if len(s) == 0 { return nil }
    var stack []*TreeNode
    for i := 0; i < len(s); {
        if s[i] != '(' && s[i] != ')' {
            start := i
            for ; i < len(s) && s[i] != '(' && s[i] != ')'; i++ { }
            v, _ := strconv.Atoi(s[start:i])
            node := &TreeNode{v, nil, nil}
            if len(stack) != 0 {
                p := stack[len(stack)-1]
                if p.Left == nil { p.Left = node } else { p.Right = node }
            }
            stack = append(stack, node)
        } else if s[i] == '(' {
            i++
        } else { // s[i] == ')'
            stack = stack[:len(stack)-1]
            i++
        }
    }
    return stack[0]
}
