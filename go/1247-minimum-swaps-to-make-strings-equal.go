func minimumSwap(s1 string, s2 string) int {
    ret := 0
    x1, x2, y1, y2 := 0, 0, 0, 0
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            if s1[i] == 'x' { x1++ } else { y1++ }
            if s2[i] == 'x' { x2++ } else { y2++ }
        }
    }
    if (x1 + x2) & 1 == 1 || (y1 + y2) & 1 == 1 { return -1 }
    ret += x1/2 + y1/2
    ret += (x1&1)<<1
    return ret
}
