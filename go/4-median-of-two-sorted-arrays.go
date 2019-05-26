func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    smaller, larger := nums1, nums2
    if len(nums1) > len(nums2) { smaller, larger = nums2, nums1 }
    sLeft, sRight := 0, len(smaller)
    halfLen := (len(nums1) + len(nums2) + 1) / 2
    for sLeft <= sRight {
        sIdx := sLeft + (sRight - sLeft) / 2
        lIdx := halfLen - sIdx // lIdx must be more than 0.
        if sIdx > 0 && smaller[sIdx-1] > larger[lIdx] {
            sRight = sIdx - 1
        } else if sIdx < len(smaller) && smaller[sIdx] < larger[lIdx-1] {
            sLeft = sIdx + 1
        } else {
            // find the sIdx that partition smaller into smaller[:sIdx], smaller[sIdx:],
            // partition larger into larger[:lIdx], larger[lIdx:]
            // Note: sIdx could be 0 or len(smaller) and lIdx could be len(larger) or 0 respectively.
            var midLeft, midRight int
            if sIdx == 0 {
                midLeft = larger[lIdx-1]
            } else if lIdx == 0 {
                midLeft = smaller[sIdx-1]
            } else {
                midLeft = max(smaller[sIdx-1], larger[lIdx-1])
            }
            if (len(nums1) + len(nums2)) % 2 == 0 {
                 if sIdx == len(smaller) {
                    midRight = larger[lIdx]
                } else if lIdx == len(larger) {
                    midRight = smaller[sIdx]
                } else {
                    midRight = min(smaller[sIdx], larger[lIdx])
                }
                return float64(midLeft + midRight) / 2.0
            } else {
                return float64(midLeft)
            }
        }
    }
    return -1
}

func max(a, b int) int {
    if a > b { return a } else { return b } 
}

func min(a, b int) int {
    if a > b { return b } else { return a } 
}
