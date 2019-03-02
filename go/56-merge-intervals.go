/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "sort"
import "fmt"

type Intervals []Interval
func (s Intervals) Len() int {return len(s)}
func (s Intervals) Swap(i, j int) {s[i], s[j] = s[j], s[i]}
func (s Intervals) Less(i, j int) bool {return s[i].Start < s[j].Start }

func merge(intervals []Interval) []Interval {
    if len(intervals) <= 1 {
        return intervals
    }
    sort.Sort(Intervals(intervals))
    var results []Interval
    init := intervals[0]
    for _, itv := range intervals[1:] {
        if init.End >= itv.Start {
             if itv.End > init.End {
                 init.End = itv.End
             }
        } else {
             results = append(results, init)
             init = itv
        }
    } 
    results = append(results, init)
    return results
}
