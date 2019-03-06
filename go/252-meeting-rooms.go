/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "sort"
func canAttendMeetings(intervals []Interval) bool {
    if len(intervals) <= 1 {
        return true
    }
    sort.Sort(Intervals(intervals))
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start < intervals[i-1].End {
            return false
        }
    }
    return true
}

type Intervals []Interval
func (itl Intervals) Less (i, j int) bool { return itl[i].Start < itl[j].Start}
func (itl Intervals) Swap (i, j int) { itl[i], itl[j] = itl[j], itl[i]}
func (itl Intervals) Len() int { return len(itl)}
