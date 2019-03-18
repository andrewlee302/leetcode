/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strings"
import "strconv"
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {return []string{}}
    var path []int
    var ret []string
    DFS(root, &path, 0, &ret)
    return ret
}

func DFS(node *TreeNode, path *[]int, idx int, ret *[]string) {
    if node == nil {
        return
    }
    if idx < len(*path) {
        (*path)[idx] = node.Val
    } else {
        *path = append(*path, node.Val)
    }
    if node.Left == nil && node.Right == nil {
        var builder strings.Builder
        for i := 0; i < idx; i++ {
            builder.WriteString(strconv.Itoa((*path)[i]))
            builder.WriteString("->")
        }
        builder.WriteString(strconv.Itoa((*path)[idx]))
        *ret = append(*ret, builder.String())
    }
    DFS(node.Left, path, idx + 1, ret)
    DFS(node.Right, path, idx + 1, ret)
}
