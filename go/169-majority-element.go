// quick-select, find median number of an array.
func majorityElement(nums []int) int {
    left, right := 0, len(nums) - 1
    mid1, mid2 := left + (right-left)/2, left + (right-left+1)/2
    for left <= right {
        p := left
        for i := left; i < right; i++ {
            if nums[i] < nums[right] {
                nums[p], nums[i] = nums[i], nums[p]
                p++
            }
        }
        nums[p], nums[right] = nums[right], nums[p]
        if p == mid1 || p == mid2 {
            return nums[mid1]
        } else if p < mid1 {
            left = p + 1
        } else { // p > mid2
            right = p - 1
        }
    }
    return -1
}

// Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
    candidate := 0
    count := 0
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if num == candidate { count++ } else { count-- }
    }
    return candidate
}
