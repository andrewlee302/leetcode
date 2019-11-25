func sortedListToBST(head *ListNode) *TreeNode {
    var nums []int
    for ; head != nil; head = head.Next {
        nums = append(nums, head.Val)
    }
    return constructBalancedBST(nums)
}

func constructBalancedBST(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    mid := (len(nums) - 1) / 2
    return &TreeNode{nums[mid], constructBalancedBST(nums[:mid]), constructBalancedBST(nums[mid+1:])}
}
