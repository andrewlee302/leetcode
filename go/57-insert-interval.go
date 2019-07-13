func insert(intervals [][]int, newInterval []int) [][]int {
    ret := make([][]int, 0, len(intervals)+1)
    i := 0
    for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ { ret = append(ret, intervals[i]) }
    if i < len(intervals) && intervals[i][0] <= newInterval[1] { 
        newInterval[0] = min(intervals[i][0], newInterval[0]) 
    }
    for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ { }
    if i-1 >= 0 && intervals[i-1][1] >= newInterval[0] {
        newInterval[1] = max(intervals[i-1][1], newInterval[1])
    }
    ret = append(ret, newInterval)
    for ; i < len(intervals); i++ { ret = append(ret, intervals[i]) }
    return ret
}

func min(i, j int) int { if i < j { return i } else { return j } }
func max(i, j int) int { if i > j { return i } else { return j } }
