/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
// recursion
func intervalIntersection(A [][]int, B [][]int) [][]int {
    var ret [][]int
    for i, j := 0, 0; i < len(A) && j < len(B); {
        itv1, itv2 := A[i], B[j]
        if itv2[0] < itv1[0] { itv1, itv2 = itv2, itv1 }
        if itv1[1] >= itv2[0] {
            ret = append(ret, []int{max(itv1[0], itv2[0]), min(itv1[1], itv2[1])})
        }
        if A[i][1] == B[j][1] {
            i, j = i+1, j+1
        } else if A[i][1] > B[j][1] {
            j++
        } else {
            i++
        }
    }
    return ret
}

func min(a, b int) int { if a < b { return a } else { return b } }
func max(a, b int) int { if a > b { return a } else { return b } }

// iteration.
func intervalIntersection(A []Interval, B []Interval) []Interval  {
    if len(A) == 0 || len(B) == 0 { return []Interval{} }
    var ret []Interval
    iA, iB := 0, 0
    for iA < len(A) && iB< len(B) {
        maxStart := A[iA].Start
        if B[iB].Start > maxStart { maxStart = B[iB].Start }
        minEnd := A[iA].End
        if B[iB].End < minEnd { minEnd = B[iB].End }
        if minEnd >= maxStart {
            ret = append(ret, Interval{maxStart, minEnd})
        }
        if A[iA].End > B[iB].End {
            iB++
        } else if A[iA].End < B[iB].End {
            iA++
        } else {
            iA, iB = iA + 1, iB + 1
        }
    }
    return ret
}
