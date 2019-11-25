func singleNonDuplicate(nums []int) int {
    i, j := 0, len(nums) - 1
    for i < j {
        mid := i + (j-i)>>1
        var l, r int
        if mid & 1 == 0 { l, r = mid, mid+1 } else { l, r = mid-1, mid }
        if r < len(nums) && nums[l] == nums[r] {
            i = mid + 1
        } else {
            j = mid
        }
    }
    return nums[i]
}

import "sort"
func singleNonDuplicate(nums []int) int {
    return nums[sort.Search(len(nums), func(i int) bool {
        var l, r int
        if i & 1 == 0 { l, r = i, i+1 } else { l, r = i-1, i }
        return r == len(nums) || nums[l] != nums[r]
    })]
}
