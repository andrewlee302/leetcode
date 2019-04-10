// recursion
func searchRange(nums []int, target int) []int {
    left := helper(nums, 0, len(nums)-1, target, true)
    right := helper(nums, 0, len(nums)-1, target, false)
    return []int{left, right}
}

func helper(nums []int, l, r, target int, left bool) int {
    if l > r { return -1 }
    if l == r { if target != nums[l] { return -1 } else { return l } }
    mid := l + (r-l)/2
    if nums[mid] > target { 
        return helper(nums, l, mid-1, target, left)
    } else if nums[mid] < target {
        return helper(nums, mid+1, r, target, left)
    } else {
        if left { return helper(nums, l, mid, target, left) } else {
            // Note, helper(nums, mid, r, target, left) may incur the endless loop. [1,1] find 1.
            if nums[mid+1] == target {
                return helper(nums, mid+1, r, target, left) 
            } else {
                return mid
            }    
        }   
    }
} 

// iteration
func searchRange(nums []int, target int) []int {
    leftP, rightP := -1, -1
    l, r := 0, len(nums) - 1
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] > target {
            r = mid - 1
        } else if nums[mid] < target {
            l = mid + 1
        } else { r = mid }
    }
    if l == r && nums[l] == target { leftP = l }
    l, r = 0, len(nums) - 1
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] > target {
            r = mid - 1
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            if nums[mid+1] == target { l = mid + 1 } else {
                rightP = mid
                break
            }
        }
    }
    if l == r && nums[l] == target { rightP = l }
    return []int{leftP, rightP}
}
