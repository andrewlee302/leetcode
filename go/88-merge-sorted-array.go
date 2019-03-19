func merge(nums1 []int, m int, nums2 []int, n int)  {
    p, i1, i2 := m + n - 1, m - 1, n - 1
    for ; i1 >= 0 && i2 >= 0; p-- {
        if nums1[i1] >= nums2[i2] {
            nums1[p] = nums1[i1]
            i1--   
        } else {
            nums1[p] = nums2[i2]
            i2--
        }
    }
    for ; i2 >= 0; p-- {
        nums1[p] = nums2[i2]
        i2--
    }
}
