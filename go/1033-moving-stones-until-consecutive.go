import "sort"
func numMovesStones(a int, b int, c int) []int {
    arr := []int{a, b, c}
    sort.Ints(arr)
    diffs := []int{arr[1] - arr[0], arr[2] - arr[1]}
    sort.Ints(diffs)
    if diffs[0] == 1 && diffs[1] == 1 { return []int{0, 0} }
    if diffs[0] == 1 { return []int{1, diffs[1] - 1} }
    if diffs[0] == 2 { return []int{1, diffs[0] + diffs[1] - 2} }
    return []int{2, diffs[0] + diffs[1] - 2}
}
