/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"
type BSTIterator struct {
    stack *list.List
}


func pushAllLeft(stack *list.List, node *TreeNode) {
    for ; node != nil; node = node.Left {
        stack.PushBack(node)
    }
}

func Constructor(root *TreeNode) BSTIterator {
    stack := list.New()
    pushAllLeft(stack, root)
    return BSTIterator{stack:stack}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    e := this.stack.Back()
    this.stack.Remove(e)
    node := e.Value.(*TreeNode)
    pushAllLeft(this.stack, node.Right)
    return node.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.stack.Len() != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
