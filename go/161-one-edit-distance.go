func isOneEditDistance(s string, t string) bool {
    diff := len(t) - len(s)
    if diff == 0 { // replace
        replaceCnt := 0
        for i := 0; i < len(s); i++ {
            if s[i] != t[i] { replaceCnt++ }
        }
        if replaceCnt == 1 { return true } else { return false }
    } else if diff == 1 { // add
        diffCnt := 0
        for i, j := 0, 0; i < len(s) && j < len(t);  {
            if s[i] != t[j] { diffCnt++ } else { i++ }
            j++
        }
        if diffCnt <= 1 { return true } else { return false }
    } else if diff == -1 {
        return isOneEditDistance(t, s)
    } else {
        return false
    }
}
