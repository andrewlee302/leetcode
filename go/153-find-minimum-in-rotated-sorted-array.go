func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        if nums[left] <= nums[right] {return nums[left]} // equals when left == right
        mid := (left+right)/2
        if nums[mid] >= nums[left] { // mid may equal left.
            left = mid + 1
        } else if nums[mid] < nums[0] {
            right = mid
        }
    }
    return -1
}
