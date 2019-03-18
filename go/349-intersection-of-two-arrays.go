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
