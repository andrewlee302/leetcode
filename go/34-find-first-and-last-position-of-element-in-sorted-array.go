// Golang sort.Search()
import "sort"
func searchRange(nums []int, target int) []int {
    upperBound := sort.Search(len(nums), func(i int) bool { return nums[i] > target })
    lowerBound := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
    if lowerBound == len(nums) || lowerBound == upperBound { return []int{-1, -1} }
    return []int{lowerBound, upperBound-1}
}

// By myself.
func searchRange(nums []int, target int) []int {
    l, r := 0, len(nums)
    for l < r {
        mid := l + (r-l) >> 1
        if nums[mid] <= target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    upperBound := l
    l, r = 0, len(nums)
    for l < r {
        mid := l + (r-l) >> 1
        if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    lowerBound := l
    if lowerBound == len(nums) || lowerBound == upperBound { return []int{-1, -1} }
    return []int{lowerBound, upperBound-1}
}
