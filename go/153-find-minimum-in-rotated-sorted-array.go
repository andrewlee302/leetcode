func findMin(nums []int) int {
    if len(nums) == 1 { return nums[0] }
    left, right := 0, len(nums)-1
    for left <= right {
        // equals when left == right, because no duplicate
        if nums[left] <= nums[right] { return nums[left] } 
        mid := (left+right)/2
        if nums[mid] >= nums[left] { // mid may equal left.
            left = mid + 1
        } else if nums[mid] < nums[right] {
            right = mid
        }
    }
    return -1
}
