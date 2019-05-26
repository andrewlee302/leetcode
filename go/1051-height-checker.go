import "sort"
func heightChecker(heights []int) int {
    cp := make([]int, len(heights))
    copy(cp, heights)
    sort.Ints(cp)
    diffCnt := 0
    for i := 0; i < len(heights); i++ {
        if heights[i] != cp[i] {
            diffCnt++
        }
    }
    return diffCnt
}
