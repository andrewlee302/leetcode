import "sort"
func findMinArrowShots(points [][]int) int {
    if len(points) == 0 { return 0 }
    sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
    arrowInterval := points[0]
    cnt := 1
    for i := 1; i < len(points); i++ {
        if points[i][0] > arrowInterval[1] { 
            arrowInterval = points[i]
            cnt++
        } else {
            arrowInterval[0] = points[i][0]
            if points[i][1] < arrowInterval[1] { arrowInterval[1] = points[i][1] }
        }
    }
    return cnt
}
