/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil { return 0 }
    max, _ := helper(root)
    return max
}

// maxPathSum(root) = max(maxPathSum(root.Left), maxPathSum(root.Right), 
// 		maxPathSumFrom(root.Left) (if>0) + maxPathSumFrom(root.Right) (if>0) + root.Val)
// maxPathSumFrom(root) = max(maxPathSumFrom(root.Left), maxPathSumFrom(root.Right)) + root.Val
// root mustn't be null
// return maximum sum, maximum sum from root
func helper(root *TreeNode) (int, int) {
    maxFromRoot, max, maxThroughRoot := root.Val, root.Val, root.Val
    if root.Left != nil {
        leftMax, leftMaxFromRoot := helper(root.Left)  
        if leftMaxFromRoot + root.Val > maxFromRoot { 
            maxFromRoot = leftMaxFromRoot + root.Val 
        }
        if leftMaxFromRoot > 0 { maxThroughRoot += leftMaxFromRoot }
        if leftMax > max { max = leftMax }
    }
    if root.Right != nil {
        rightMax, rightMaxFromRoot := helper(root.Right)  
        if rightMaxFromRoot + root.Val > maxFromRoot { 
            maxFromRoot = rightMaxFromRoot + root.Val
        }
        if rightMaxFromRoot > 0 { maxThroughRoot += rightMaxFromRoot }
        if rightMax > max { max = rightMax }
    }
    if maxThroughRoot > max { max = maxThroughRoot }
    return max, maxFromRoot
}

