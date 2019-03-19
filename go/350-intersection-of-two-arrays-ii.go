func intersect(nums1 []int, nums2 []int) []int {
    m1, m2 := make(map[int]int), make(map[int]int)
    for _, v := range nums1 { m1[v]++ }
    for _, v := range nums2 { m2[v]++ }
    if len(m1) < len(m2) { 
        return mapIntersect(m1, m2) 
    } else {
        return mapIntersect(m2, m1) 
    }
}

func mapIntersect(m1, m2 map[int]int) []int {
    var ret []int
    for k1, v1 := range m1 {
        if times, ok := m2[k1]; ok {
            if times > v1 { times = v1 }
            for i := 0; i < times; i++ {
                ret = append(ret, k1)
            }
        }  
    }
    return ret
}

// assumed sorted.
import "sort"
func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) == 0 || len(nums2) == 0 { return []int{} }
    sort.Sort(sort.IntSlice(nums1))
    sort.Sort(sort.IntSlice(nums2))
    var ret []int
    for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
        if nums1[i] == nums2[j] {
            ret = append(ret, nums1[i])
            i++
            j++
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }
    return ret
}
