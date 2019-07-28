func numEquivDominoPairs(dominoes [][]int) int {
    m := make(map[string]int)
    for _, d := range dominoes {
        if d[1] < d[0] { d[0], d[1] = d[1], d[0] }
        m[strconv.Itoa(d[0]) + "," + strconv.Itoa(d[1])]++
    }
    total := 0
    for _, cnt := range m {
        total += cnt * (cnt-1) / 2
    }
    return total
}
