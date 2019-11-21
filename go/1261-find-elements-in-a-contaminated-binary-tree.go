/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    root *TreeNode
    maxV int
    height int
}


func Constructor(root *TreeNode) FindElements {
    if root == nil { return FindElements{nil, -1, 0} }
    root.Val = 0
    maxV := 0
    DFS(root, &maxV)
    temp := maxV + 1
    height := 0
    for ; temp >= 1; height++ { temp >>= 1 }
    return FindElements{root, maxV, height}
}

func DFS(root *TreeNode, maxV *int) {
    if *maxV < root.Val { *maxV = root.Val }
    if root.Left != nil { 
        root.Left.Val = root.Val * 2 + 1 
        DFS(root.Left, maxV)
    }
    if root.Right != nil { 
        root.Right.Val = root.Val * 2 + 2
        DFS(root.Right, maxV)
    }
} 


func (this *FindElements) Find(target int) bool {
    if target < 0 || target > this.maxV { return false }
    temp := target + 1
    height := 0
    for ; temp >= 1; height++ { temp >>= 1 }
    offset := target + 1 - (1<<uint(height-1))
    l, r := 0, 1<<uint(height-1) - 1
    node := this.root
    for i := 0; i < height; i++ {
        if node == nil { return false }
        mid := l + (r-l)/2
        if offset > mid {
            node = node.Right
            l = mid+1
        } else {
            node = node.Left
            r = mid
        }
    }
    return true
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
