import "sort"
func videoStitching(clips [][]int, T int) int {
    if T == 0 { return 0 }
    sort.Slice(clips, func(i, j int) bool { return clips[i][0] < clips[j][0] })
    cnt := 0
    intervalEnd := 0
    startIdx := 0 
    for {
        maxEnd := intervalEnd
        i := startIdx
        for ; i < len(clips) && clips[i][0] <= intervalEnd; i++ {
            if clips[i][1] > maxEnd {
                maxEnd = clips[i][1]
            }
        }
        startIdx = i
        cnt++
        if maxEnd == intervalEnd { return -1 }
        intervalEnd = maxEnd
        if intervalEnd >= T { return cnt }
    }
    return -1
}
