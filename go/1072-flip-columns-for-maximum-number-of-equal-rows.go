func maxEqualRowsAfterFlips(matrix [][]int) int {
    m := make(map[string]int)
    for _, a := range matrix {
        if a[0] == 1 { for i := 0; i < len(a); i++ { a[i] = 1-a[i] } }
        m[str(a)]++
    }
    ret := 0
    for _, cnt := range m { ret = max(ret, cnt) }
    return ret
}

func str(a []int) string {
    bs := make([]byte, len(a))
    for i, v := range a { bs[i] = byte(v) + '0' }
    return string(bs)
}

func max(i ,j int) int { if i > j { return i } else { return j } }
