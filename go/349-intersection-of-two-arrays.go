func intersection(nums1 []int, nums2 []int) []int {
    set1, set2 := make(map[int]bool), make(map[int]bool)
    for _, v := range nums1 { set1[v] = true }
    for _, v := range nums2 { set2[v] = true }
    if len(set1) < len(set2) { 
        return setIntersection(set1, set2) 
    } else { 
        return setIntersection(set2, set1) 
    }
}

func setIntersection(set1, set2 map[int]bool) []int {
    var ret []int
    for v, _ := range set1 {
        if set2[v] {
            ret = append(ret, v)
        }
    }
    return ret
}

// if sorted, how to solve.
import "sort"
func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1) == 0 || len(nums2) == 0 { return []int{} }
    sort.Sort(sort.IntSlice(nums1))
    sort.Sort(sort.IntSlice(nums2))
    leftArray, rightArray := nums1, nums2
    if nums1[0] > nums2[0] {
        leftArray, rightArray = nums2, nums1
    }
    var ret []int
    i, j := 0, 0 // leftArray index, rightArray index
    for ; i < len(leftArray) && j < len(rightArray); {
        for ; i < len(leftArray) && leftArray[i] < rightArray[j]; i++ { }
        if i < len(leftArray) && leftArray[i] == rightArray[j] { 
            ret = append(ret, leftArray[i]) 
            for ; i < len(leftArray) && leftArray[i] == rightArray[j]; i++ { }
        }
        if i < len(leftArray) {
            for ; j < len(rightArray) && rightArray[j] < leftArray[i]; j++ { }
            if j < len(rightArray) && rightArray[j] == leftArray[i] { 
                ret = append(ret, rightArray[j]) 
                for ; j < len(rightArray) && rightArray[j] == leftArray[i]; j++ { }
            }
        }
    }
    return ret
}


