func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    minIdx := 0
    minV := nums[0]
    for i, v := range nums {
        if v < minV {
            minIdx = i
            minV = v
        }
    }
    if minIdx == 0 {
        return binarySearch(nums, 0, len(nums)-1, target)
    }
    if target <= nums[len(nums)-1] {
        return binarySearch(nums, minIdx, len(nums)-1, target)
    } else {
        return binarySearch(nums, 0, minIdx-1, target)
    }
    
}

func binarySearch(nums []int, left, right int, target int) int {
    if left == right {
        if nums[left] == target {
            return left
        } else {
            return -1
        }
    }
    mid := (left+right)/2
    if target <= nums[mid] {
        return binarySearch(nums, left, mid , target)
    } else {
        return binarySearch(nums, mid+1, right, target)
    }
}
