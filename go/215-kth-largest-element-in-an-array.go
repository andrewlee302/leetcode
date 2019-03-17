func findKthLargest(nums []int, k int) int {
    quickselect(nums, 0, len(nums)-1, k)
    return nums[k-1]
}

func quickselect(nums []int, left, right, k int) {
    if left >= right {
        return
    }
    pivotValue := nums[right]
    storeIndex := left
    for i := left; i < right; i++ {
        if nums[i] >  pivotValue {
            nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
            storeIndex++
        }
    }
    nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
    if k-1 == storeIndex {
        return
    } else if k-1 < storeIndex {
        quickselect(nums, left, storeIndex-1, k)
    } else {
        quickselect(nums, storeIndex+1, right, k)
    }
}
