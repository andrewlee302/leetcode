func missingElement(nums []int, k int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := l + (r-l) >> 1
        if missNum := (nums[mid] - nums[0]) - (mid-0); missNum < k {
            l = mid + 1
        } else {
            r = mid
        }
    }
    // missNumBeforeL must be non-negative.
    missNumBeforeL := (nums[l-1] - nums[0]) - (l-1)
    return nums[l-1] + (k - missNumBeforeL)
}

// sort.Search()
import "sort"
func missingElement(nums []int, k int) int {
    pos := sort.Search(len(nums), func(i int) bool {
        missNum := (nums[i] - nums[0]) - (i-0)
        return missNum >= k
    })
    // missNumOnPosMinus1 must be non-negative.
    missNumOnPosMinus1 := (nums[pos-1] - nums[0]) - (pos-1)
    return nums[pos-1] + (k - missNumOnPosMinus1)
}
