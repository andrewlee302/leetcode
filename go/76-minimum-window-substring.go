func minWindow(s string, t string) string {
    if len(s) < len(t) { return "" }
    tDict := make(map[byte]int)
    for i := 0; i < len(t); i++ { tDict[t[i]]++ }
    tDictSize := len(tDict)
    subStrDictMatchCnt := 0
    left, right := 0, 0
    minimalLen := len(s) + 1
    ret := ""
    for right < len(s) {
        if remaining, ok := tDict[s[right]]; ok {
            remaining--
            tDict[s[right]] = remaining
            if remaining == 0 { subStrDictMatchCnt++ }
        }
        for subStrDictMatchCnt == tDictSize {
            // left mustn't be more than right
            if remaining, ok := tDict[s[left]]; ok {
                tDict[s[left]] = remaining + 1
                if remaining == 0 { 
                    subStrDictMatchCnt-- 
                    // left ~ right
                    if right-left+1 < minimalLen {
                        minimalLen = right-left+1
                        ret = s[left: right+1]
                    }
                }
            }
            left++
        }
        right++
    }
    return ret
}

