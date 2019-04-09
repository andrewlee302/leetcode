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
    sort.Ints(nums1)
    sort.Ints(nums2)
    var ret []int
    i1, i2 := 0, 0
    for i1 < len(nums1) && i2 < len(nums2) {
        if nums1[i1] > nums2[i2] {
            for i2++; i2 < len(nums2) && nums2[i2] == nums2[i2-1]; i2++ { }
        } else if nums1[i1] < nums2[i2] {
            for i1++; i1 < len(nums1) && nums1[i1] == nums1[i1-1]; i1++ { }
        } else {
            ret = append(ret, nums1[i1])
            for i1++; i1 < len(nums1) && nums1[i1] == nums1[i1-1]; i1++ { }
            for i2++; i2 < len(nums2) && nums2[i2] == nums2[i2-1]; i2++ { }
        }
    }
    return ret
}
