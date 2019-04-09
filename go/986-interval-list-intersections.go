/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
 // recursion
func intervalIntersection(A []Interval, B []Interval) []Interval  {
    if len(A) == 0 || len(B) == 0 { return []Interval{} }
    if A[0].Start > B[0].Start { return intervalIntersection(B, A) }
  
    // Here, A[0].Start <= B[0].Start
    if A[0].End >= B[0].Start { 
        var overlap Interval
        var laterRet []Interval
        if B[0].End < A[0].End {
            overlap = Interval{B[0].Start, B[0].End }
            laterRet = intervalIntersection(A, B[1:])
        } else {
            overlap = Interval{B[0].Start, A[0].End }
            laterRet = intervalIntersection(A[1:], B)
        }
        ret := make([]Interval, len(laterRet)+1)
        ret[0] = overlap
        copy(ret[1:], laterRet)
        return ret
    } else {
        return intervalIntersection(A[1:], B)
    }
}

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
