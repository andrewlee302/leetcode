// binary search.
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, mountainArr *MountainArray) int {
    return helper(target, mountainArr, 0, mountainArr.length()-1)
}

func helper(target int, mountainArr *MountainArray, l, r int) int {
    if l > r { return -1 }
    mid := l + (r-l)/2
    midV := mountainArr.get(mid)
    lVal, rVal := -1, -1
    if mid != l { lVal = mountainArr.get(mid-1) }
    if mid != r { rVal = mountainArr.get(mid+1) }
    if midV < target {
        if midV > lVal && midV > rVal { return -1 }
        if midV > lVal { 
            return helper(target, mountainArr, mid+1, r)
        } else { 
            return helper(target, mountainArr, l, mid-1)
        }
    } else if midV > target {
        lRet, rRet := -1, -1
        lRet = helper(target, mountainArr, l, mid-1)
        if lRet != -1 { return lRet }
        rRet = helper(target, mountainArr, mid+1, r)
        return rRet
    } else {
        if midV > lVal && midV > rVal { return mid }
        if midV > lVal { 
            return mid
        } else { 
            lRet := helper(target, mountainArr, l, mid-1)
            if lRet != -1 { return lRet } else { return mid }
        }
    }
}
