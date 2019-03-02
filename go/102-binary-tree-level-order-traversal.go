/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"

type TreeNodeLevel struct {
    *TreeNode
    Level int
}
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    return BFS(root)
}

func BFS(root *TreeNode) [][]int {
    var results [][]int
    queue := list.New()
    queue.PushBack(&TreeNodeLevel{root, 0})
    for e := queue.Front(); e != nil; e = e.Next() {
        node := e.Value.(*TreeNodeLevel)
        if node.Level > len(results) - 1 {
            results = append(results, []int{node.Val})
        } else {
            results[node.Level] = append(results[node.Level], node.Val)
        }
        if node.Left != nil {
            queue.PushBack(&TreeNodeLevel{node.Left, node.Level+1})
        }
        if node.Right != nil {
            queue.PushBack(&TreeNodeLevel{node.Right, node.Level+1})
        }
    }
    return results
}
