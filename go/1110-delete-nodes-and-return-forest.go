// Post-order tarversal.
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    var ret []*TreeNode
    tag := make([]bool, 1001)
    for _, delNum := range to_delete {
        tag[delNum] = true
    }
    DFS(&root, tag, &ret)
    if root != nil { ret = append(ret, root) }
    return ret
}

func DFS(rootPtr **TreeNode, tag []bool, ret *[]*TreeNode) {
    if *rootPtr == nil { return } 
    root := *rootPtr
    DFS(&(root.Left), tag, ret)
    DFS(&(root.Right), tag, ret)
    if tag[root.Val] {
        if root.Left != nil { *ret = append(*ret, root.Left) }
        if root.Right != nil { *ret = append(*ret, root.Right) }
        root.Left, root.Right = nil, nil
        *rootPtr = nil
    }
}
