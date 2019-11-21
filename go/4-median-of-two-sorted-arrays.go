// binary search.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) { nums1, nums2 = nums2, nums1 }
    m, n := len(nums1), len(nums2)
    L := (m+n+1)/2
    p1, p2 := 0, len(nums1)
    for p1 <= p2 { // must find it.
        i := p1 + (p2-p1)/2
        j := L - i // j >= 0
        if i > 0 && nums1[i-1] > nums2[j] {
            p2 = i - 1
        } else if i < len(nums1) && nums1[i] < nums2[j-1] { // j > 0
            p1 = i + 1
        } else { // find it
            var LMax int
            if i == 0 {
                LMax = nums2[j-1]
            } else if j == 0 {
                LMax = nums1[i-1]
            } else {
                LMax = max(nums1[i-1], nums2[j-1])
            }
            if (m + n) % 2 == 0 {
                // RMax should be considered here as there is a corner case: [],[1] 
                var RMax int
                if i == len(nums1) {
                    RMax = nums2[j]
                } else if j == len(nums2) {
                    RMax = nums1[i]
                } else {
                    RMax = min(nums1[i], nums2[j])
                }
                return float64(LMax + RMax) / 2.0
            } else {
                return float64(LMax)
            }
        }
    }
    return -1
}

func min(i, j int) int { if i < j { return i } else { return j } }
func max(i, j int) int { if i > j { return i } else { return j } }
