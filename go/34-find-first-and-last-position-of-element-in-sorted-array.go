func searchRange(nums []int, target int) []int {
    return binarySelect(nums, 0, len(nums) - 1, target)
}

func binarySelect(nums[]int, left, right int, target int) []int {
    if left > right { return []int{-1, -1} }
    mid := left + (right - left) / 2
    if target == nums[mid] {
        m, n := mid, mid
        for ; m >= left && nums[m] == target; m-- { }
        for ; n <= right && nums[n] == target; n++ { }
        return []int{m+1, n-1}
    } else if target < nums[mid] { 
        return binarySelect(nums, left, mid - 1, target) 
    } else { return binarySelect(nums, mid + 1, right, target) }
}
