/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    if K == 0 { return []int{target.Val} }
    graph := make(map[int][]int)
    DFS(root, nil, graph)
    visited := make(map[int]bool)
    queue := list.New()
    queue.PushBack(target.Val)
    visited[target.Val] = true
    dist := 0
    for queue.Len() != 0 {
        size := queue.Len() 
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            v := e.Value.(int)
            for _, next := range graph[v] {
                if !visited[next] {
                    visited[next] = true
                    queue.PushBack(next)
                }
            }
        }
        dist += 1
        if dist == K { break }
    }
    ret := make([]int, 0, queue.Len())
    for e := queue.Front(); e != nil; e = e.Next() {
        ret = append(ret, e.Value.(int))
    }
    return ret
}

func DFS(root, parent *TreeNode, graph map[int][]int) {
    if root == nil { return }
    if parent != nil {
        graph[root.Val] = append(graph[root.Val], parent.Val)
        graph[parent.Val] = append(graph[parent.Val], root.Val)
    }
    DFS(root.Left, root, graph)
    DFS(root.Right, root, graph)
}
