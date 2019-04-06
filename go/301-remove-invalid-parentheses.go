func removeInvalidParentheses(s string) []string {
    leftRm, rightRm := 0, 0
    matchDiff := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            matchDiff++
        } else if s[i] == ')' {
            matchDiff--
            if matchDiff < 0 {
                rightRm++
                matchDiff = 0
            }
        }
    }
    leftRm = matchDiff
    buf := make([]byte, len(s))
    m := make(map[string]bool)
    helper(s, buf, 0, 0, 0, leftRm, rightRm, m)
    var ret []string
    for str, _ := range m {
        ret = append(ret, str)
    }
    return ret
}

// matchDiff is always zero at the begin of the invocation.
func helper(s string, buf []byte, bufCnt, idx, matchDiff, leftRm, rightRm int, m map[string]bool) {
    if idx == len(buf) {
        if matchDiff == 0 && leftRm == 0 && rightRm == 0 {
            m[string(buf[:bufCnt])] = true
        }
        return
    }
    // try to discard
    if s[idx] == '(' && leftRm > 0 {
        helper(s, buf, bufCnt, idx+1, matchDiff, leftRm-1, rightRm, m) 
    } else if s[idx] == ')' && rightRm > 0 {
        helper(s, buf, bufCnt, idx+1, matchDiff, leftRm, rightRm-1, m) 
    }
    
    // try to keep
    buf[bufCnt] = s[idx]
    bufCnt++
    if s[idx] == '(' {
        helper(s, buf, bufCnt, idx+1, matchDiff+1, leftRm, rightRm, m)
    } else if s[idx] == ')' {
        if matchDiff - 1 >= 0 { helper(s, buf, bufCnt, idx+1, matchDiff-1, leftRm, rightRm, m) }
    } else {
        helper(s, buf, bufCnt, idx+1, matchDiff, leftRm, rightRm, m)
    }
}
