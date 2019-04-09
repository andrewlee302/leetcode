func search(nums []int, target int) int {
    if len(nums) == 0 { return -1 }
    rotateIdx := findMinimalIndex(nums, 0, len(nums)-1)
    if rotateIdx == 0 { return binarySearch(nums, 0, len(nums)-1, target) }
    if nums[0] <= target { 
        return binarySearch(nums, 0, rotateIdx-1, target)
    } else {
        return binarySearch(nums, rotateIdx, len(nums)-1, target)
    } 
}

func findMinimalIndex(nums []int, l, r int) int {
    if l == r || nums[l] < nums[r] { return l }
    // 2~ items
    mid := l + (r - l) / 2
    if nums[mid] > nums[mid+1] { return mid+1 }
    if nums[mid] > nums[l] { return findMinimalIndex(nums, mid+1, r) } else {
        return findMinimalIndex(nums, l, mid)
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
