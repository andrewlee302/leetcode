func removeInvalidParentheses(s string) []string {
    diff := 0
    rmLeftCnt, rmRightCnt := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' { diff++ } else if s[i] == ')' { diff-- }
        if diff < 0 {
            diff = 0
            rmRightCnt++
        }
    }
    if diff > 0 { rmLeftCnt = diff }
    retMap := make(map[string]bool)
    buf := make([]byte, len(s))
    helper(s, 0, buf, 0, 0, rmLeftCnt, rmRightCnt, retMap)
    var ret []string
    for val, _ := range retMap { ret = append(ret, val) }
    return ret
}

// diff is always non-negative.
func helper(s string, idx int, buf []byte, bufIdx int, diff, rmLeftCnt, rmRightCnt int, ret map[string]bool) {
    if idx == len(s) {
        if rmLeftCnt == 0 && rmRightCnt == 0 { ret[string(buf[:bufIdx])] = true }
        return
    }
    if s[idx] == '(' && rmLeftCnt > 0 {
        helper(s, idx+1, buf, bufIdx, diff, rmLeftCnt-1, rmRightCnt, ret)
    } else if s[idx] == ')' && rmRightCnt > 0 {
        helper(s, idx+1, buf, bufIdx, diff, rmLeftCnt, rmRightCnt-1, ret)
    }

    if !(s[idx] == ')' && diff == 0) { // keep
        if s[idx] == '(' { diff++ } else if s[idx] == ')' { diff-- }
        buf[bufIdx] = s[idx]
        helper(s, idx+1, buf, bufIdx+1, diff, rmLeftCnt, rmRightCnt, ret)
    }
}
