/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"
type Wrap struct {
    node *TreeNode
    depth int
}

func recoverFromPreorder(S string) *TreeNode {
    var head *TreeNode
    var stack []Wrap
    for i := 0; i < len(S); {
        start := i
        for ; i < len(S) && S[i] == '-'; i++ { }
        depth := i - start
        start = i
        for ; i < len(S) && S[i] != '-'; i++ { }
        val, _ := strconv.Atoi(S[start:i])
        node := &TreeNode{Val: val}
        if head == nil { 
            head = node 
            stack = append(stack, Wrap{node, depth})
        } else {
            for len(stack) != 0 {
                top := stack[len(stack)-1]
                if top.depth == depth - 1 {
                    if top.node.Left == nil {
                        top.node.Left = node
                    } else { // right must be nil
                        top.node.Right = node
                    }
                    stack = append(stack, Wrap{node, depth})
                    break
                }
                stack = stack[:len(stack)-1]
            }
        }
    }
    return head
}
