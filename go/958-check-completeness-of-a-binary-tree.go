/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
func isCompleteTree(root *TreeNode) bool {
    queue := list.New()
    queue.PushBack(root)
    expectedSize := 1
    isFull := true
    isCompact := true
    for queue.Len() != 0 {
        size := queue.Len()
        if size != expectedSize { 
            if isFull { 
                if !isCompact { return false }
                isFull = false
            } else { return false }
        }
        isNull := false
        for i := 0; i < size; i++ {
            e := queue.Front()
            queue.Remove(e)
            node := e.Value.(*TreeNode)
            if node.Left != nil { 
                if isNull { isCompact = false }
                queue.PushBack(node.Left) 
            } else { isNull = true }
            if node.Right != nil { 
                if isNull { isCompact = false }
                queue.PushBack(node.Right) 
            } else { isNull = true }
        }
        expectedSize *= 2
    }
    return true
}
